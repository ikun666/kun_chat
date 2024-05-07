package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"strconv"

	"github.com/gorilla/websocket"
	"github.com/ikun666/kun_chat/chat/api/internal/global"
	"github.com/ikun666/kun_chat/chat/api/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	conn   *websocket.Conn
}

func NewChatLogic(ctx context.Context, svcCtx *svc.ServiceContext, conn *websocket.Conn) *ChatLogic {
	return &ChatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		conn:   conn,
	}
}

func (l *ChatLogic) Chat(userId int64) {
	// todo: add your logic here and delete this line
	//获取socket连接,构造消息节点
	node := &global.Node{
		Conn:      l.conn,
		DataQueue: make(chan []byte, 50),
	}

	// 用户关系
	// 将userId和Node绑定
	global.RWLocker.Lock()
	global.ClientMap[userId] = node
	global.RWLocker.Unlock()

	// fmt.Println("uid", userId)

	// 发送消息
	go sendProc(node)
	// 接收消息
	go recProc(node)
	sendMsg(userId, []byte("欢迎进入聊天系统"))

}
func sendProc(node *global.Node) {
	for {
		select {
		case data := <-node.DataQueue:
			err := node.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				fmt.Println("写入消息失败", err)
				return
			}
			// fmt.Println("数据发送socket成功")
		}
	}
}

func recProc(node *global.Node) {
	for {
		//获取信息
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			fmt.Println("读取消息失败", err)
			return
		}

		brodMsg(data)

	}
}

func brodMsg(data []byte) {
	global.UDPChan <- data
}

func init() {
	go UdpSendProc()
	go UpdRecProc()
}

// UdpSendProc 完成upd数据发送
func UdpSendProc() {
	udpConn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		//192.168.31.147
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 3000,
		Zone: "",
	})
	if err != nil {
		fmt.Println("拨号udp端口失败", err)
		return
	}

	defer udpConn.Close()

	for {
		select {
		case data := <-global.UDPChan:
			_, err := udpConn.Write(data)
			if err != nil {
				fmt.Println("写入udp消息失败", err)
				return
			}
			// fmt.Println("数据成功发送到udp服务端:", string(data))
		}
	}

}

// UpdRecProc 完成udp数据的接收
func UpdRecProc() {
	udpConn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 3000,
	})
	if err != nil {
		fmt.Println("监听udp端口失败", err)
		return
	}

	defer udpConn.Close()

	for {
		var buf [1024]byte
		n, err := udpConn.Read(buf[0:])
		if err != nil {
			fmt.Println("读取udp数据失败", err)
			return
		}

		// fmt.Println("udp服务端接收udp数据", buf[0:n])

		//处理发送逻辑
		dispatch(buf[0:n])
	}
}

func dispatch(data []byte) {
	//解析消息
	msg := global.Message{}
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
	tids, err := global.RelationModel.GetFriends(context.TODO(), target, 2)
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
func sendMsg(id int64, msg []byte) {
	global.RWLocker.Lock()
	node, ok := global.ClientMap[id]
	global.RWLocker.Unlock()

	if !ok {
		fmt.Println("userID没有对应的node")
		return
	}

	// fmt.Println("targetID:", id, "node:", node)
	if ok {
		node.DataQueue <- msg
	}
}

// sendMsgTest 发送消息 并存储聊天记录到redis
func sendMsgAndSave(uid, tid int64, msg []byte) {

	global.RWLocker.Lock()
	node, ok := global.ClientMap[tid]
	global.RWLocker.Unlock()

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
	if uid < 1000000000 {
		key = "gmsg_" + userIdStr
	} else {
		if uid < tid {
			key = "msg_" + userIdStr + "_" + targetIdStr
		} else {
			key = "msg_" + targetIdStr + "_" + userIdStr
		}

	}
	//创建记录
	num, err := global.Redis.Zcard(key)
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(num)
	_, e := global.Redis.Zadd(key, int64(num)+1, string(msg)) //jsonMsg
	//res, e := utils.Red.Do(ctx, "zadd", key, 1, jsonMsg).Result() //备用 后续拓展 记录完整msg
	if e != nil {
		fmt.Println(e)
		return
	}

	// 设置ZSET的过期时间为10天
	// expirationTime := 24 * time.Hour * 7
	expirationTime := 7 * 24 * 60 * 60
	expireErr := global.Redis.Expire(key, expirationTime)
	if expireErr != nil {
		fmt.Println(expireErr)
		return
	}

	// fmt.Println("ZSET的过期时间已设置为10天")
	// fmt.Println(ress, expireErr)

	//将key放入全局并
	// addKeyToSaveKey(key)
}
