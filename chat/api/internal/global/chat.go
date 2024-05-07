package global

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/ikun666/kun_chat/chat/model"
)

func dispatch(data []byte) {
	//解析消息
	msg := Message{}
	err := json.Unmarshal(data, &msg)
	if err != nil {
		fmt.Println("消息解析失败", err)
		return
	}

	// fmt.Println("解析数据:", "msg.FormId", msg.FormId, "targetId:", msg.TargetId, "type:", msg.Type)
	//fmt.Println("消息类型：", msg.Type, msg.Content)
	//判断消息类型
	switch msg.Type {
	case 1: //私聊
		sendMsgAndSave(msg.FormId, msg.TargetId, data)
	case 2: //群发
		sendGroupMsg(msg.FormId, msg.TargetId, data)
	}
}

// sendGroupMsg 群发
func sendGroupMsg(formId, target int64, data []byte) (int, error) {
	//群发的逻辑：1获取到群里所有用户，然后向除开自己的每一位用户发送消息
	tids, err := RelationModel.GetFriends(context.TODO(), target, 2)
	if err != nil {
		return -1, err
	}
	// fmt.Println("群成员：", tids)
	for _, tid := range tids {
		if formId != tid {
			sendMsgAndSave(target, tid, data)
		}
	}
	return 0, nil
}

// sendMs ws向用户发送消息
// func sendMsg(id int64, msg []byte) {
// 	RWLocker.Lock()
// 	node, ok := ClientMap[id]
// 	RWLocker.Unlock()

// 	if !ok {
// 		fmt.Println("userID没有对应的node")
// 		return
// 	}

// 	// fmt.Println("targetID:", id, "node:", node)
// 	if ok {
// 		node.DataQueue <- msg
// 	}
// }

// sendMsgTest 发送消息 并存储聊天记录到redis
func sendMsgAndSave(uid, tid int64, msg []byte) {

	RWLocker.Lock()
	node, ok := ClientMap[tid]
	RWLocker.Unlock()

	// jsonMsg := Message{}
	// json.Unmarshal(msg, &jsonMsg)
	// ctx := context.Background()
	targetIdStr := strconv.Itoa(int(tid))
	userIdStr := strconv.Itoa(int(uid))

	//如果在线，需要即时推送
	if ok {
		node.DataQueue <- msg
	}

	//拼接记录名称
	var key string
	if uid < groupLimit {
		key = "msg_g_" + userIdStr
	} else {
		if uid < tid {
			key = "msg_f_" + userIdStr + "_" + targetIdStr
		} else {
			key = "msg_f_" + targetIdStr + "_" + userIdStr
		}
	}
	//创建记录
	num, err := Redis.Zcard(key)
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(num)

	_, e := Redis.Zadd(key, int64(num)+1, string(msg)) //jsonMsg
	//res, e := utils.Red.Do(ctx, "zadd", key, 1, jsonMsg).Result() //备用 后续拓展 记录完整msg
	if e != nil {
		fmt.Println(e)
		return
	}

	// 设置ZSET的过期时间为10天
	// expirationTime := 24 * time.Hour * 7
	// expirationTime := 7 * 24 * 60 * 60
	expireErr := Redis.Expire(key, expirationTime)
	if expireErr != nil {
		fmt.Println(expireErr)
		return
	}

	// fmt.Println("ZSET的过期时间已设置为10天")
	// fmt.Println(ress, expireErr)

	//将key放入全局
	addKeyToSaveKey(key)
}

// RecordPersistence 聊天记录持久化
func RecordPersistence() {

	//启动定时任务，每隔24小时执行一次
	// interval := 5 * time.Minute
	ticker := time.NewTicker(interval * time.Hour)
	defer ticker.Stop()

	// 创建一个通道用于控制并发，控制10个并发
	// concurrency := make(chan struct{}, 10)

	fmt.Println("----------持久化开始------------")

	for range ticker.C {
		fmt.Println("持久化中")
		// 处理数据持久化
		msgKeys := getSaveKey()
		fmt.Println(msgKeys)
		go func(keys []string) {

			// // 控制并发，10并发缓冲区满，G陷入阻塞
			// concurrency <- struct{}{}
			// defer func() {

			// 	//正常的G完成后，消费缓冲区
			// 	<-concurrency
			// }()

			for _, key := range keys {
				chatRecords, err := Redis.Zrange(key, 0, -1)
				if err != nil {
					fmt.Println("从Redis获取聊天记录失败:", err)
					return
				}

				fmt.Println("----------持久化中------------")
				//存储需要持久化的记录
				msg := make([]model.Message, 0, len(chatRecords))

				//持久化失败的记录归还
				backMsg := make([]string, 0, len(chatRecords))

				//开启事务
				// tx := db.DB.Begin()

				for _, record := range chatRecords {
					m := model.Message{}
					err := json.Unmarshal([]byte(record), &m)
					if err != nil {
						fmt.Println("Unmarshal fail", err.Error())

						//回滚事务
						// tx.Rollback()
						return
					}
					msg = append(msg, m)
					backMsg = append(backMsg, record)
				}

				// if err := tx.Table("messages").Save(msg).Error; err != nil {
				// 	zap.S().Info("持久化失败", err.Error())

				// 	//回滚事务
				// 	tx.Rollback()
				// 	return
				// }
				MessagesModel.BulkInsert(msg)

				if _, err := Redis.Zrem(key, backMsg); err != nil {
					fmt.Println("从Redis删除聊天记录失败:", err)

					//回滚事务
					// tx.Rollback()
					return
				}

				// //提交事务
				// tx.Commit()

				//持久化成功后移除对应key
				removeKeyFromSaveKey(key)
				fmt.Println("----------持久化成功------------")
			}
		}(msgKeys)
	}

}

// 向 saveKey 中添加 key 的函数，保护了 saveKey 的并发访问
func addKeyToSaveKey(key string) {
	saveKeyMutex.Lock()
	defer saveKeyMutex.Unlock()
	saveKey[key] = struct{}{}
}

// 获取当前的 saveKey 副本的函数，保护了 saveKey 的并发访问
func getSaveKey() []string {
	saveKeyMutex.RLock()
	defer saveKeyMutex.RUnlock()
	keys := make([]string, 0, len(saveKey))
	for key := range saveKey {
		keys = append(keys, key)
	}
	return keys
}

// 从 saveKey 中删除 key 的函数，保护了 saveKey 的并发访问
func removeKeyFromSaveKey(key string) {
	saveKeyMutex.Lock()
	defer saveKeyMutex.Unlock()
	delete(saveKey, key)
}
func SetSaveKeyFromRedis() {

	keys, cursor, err := Redis.Scan(0, "msg_*", 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(keys, cursor)
	saveKeyMutex.Lock()
	defer saveKeyMutex.Unlock()
	for i := range keys {
		saveKey[keys[i]] = struct{}{}
	}

}

// // 向 saveMsgId 中添加 key 的函数，保护了 saveMsgId 的并发访问
// func addKeyToSaveMsgId(key int64) {
// 	saveMsgMutex.Lock()
// 	defer saveMsgMutex.Unlock()
// 	saveMsgId[key] = struct{}{}
// }

// // 向 saveMsgId 中添加 key 的函数，保护了 saveMsgId 的并发访问
// func getKeyFromSaveMsgId(key int64) bool {
// 	saveMsgMutex.RLock()
// 	defer saveMsgMutex.RUnlock()
// 	_, ok := saveMsgId[key]
// 	return ok
// }
