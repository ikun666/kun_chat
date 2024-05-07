package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

// 记录用户聊天key
// var saveKeyMutex sync.Mutex
// var saveKey []string

type Message struct {
	Id         int64     `json:"id"`
	FormId     int64     `json:"form_id"`   // 信息发送者
	TargetId   int64     `json:"target_id"` // 信息接收者
	Type       int64     `json:"type"`      // 聊天类型:私聊 1 群聊2  广播3
	Media      int64     `json:"media"`     // 信息类型:文字1 图片2 音频3
	Content    string    `json:"content"`   // 消息内容
	Pic        string    `json:"pic"`       // 图片相关
	Url        string    `json:"url"`       // 文件相关
	Desc       string    `json:"desc"`      // 文件描述
	Amount     int64     `json:"amount"`    // 其他数据大小
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

// Node 构造连接
type Node struct {
	Conn      *websocket.Conn //连接
	Addr      string          //客户端地址
	DataQueue chan []byte     //消息
	// GroupSets set.Interface   //好友 / 群
}

// 映射关系
var clientMap map[int64]*Node = make(map[int64]*Node, 0)

// 读写锁
var rwLocker sync.RWMutex

// Chat	需要 ：发送者ID ，接受者ID ，消息类型，发送的内容，发送类型
func Chat(w http.ResponseWriter, r *http.Request) {
	//1.  获取参数 并 检验 token 等合法性
	query := r.URL.Query()
	fmt.Println("handle:", query)
	Id := query.Get("uid")
	//token := query.Get("token")

	userId, err := strconv.ParseInt(Id, 10, 64)
	if err != nil {
		zap.S().Info("类型转换失败", err)
		return
	}

	//升级为socket

	conn, err := (&websocket.Upgrader{
		//token 校验
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}).Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	//获取socket连接,构造消息节点
	node := &Node{
		Conn:      conn,
		DataQueue: make(chan []byte, 50),
	}

	//用户关系
	//将userId和Node绑定
	rwLocker.Lock()
	clientMap[userId] = node
	rwLocker.Unlock()

	fmt.Println("uid", userId)

	//发送消息
	go sendProc(node)
	//接收消息
	go recProc(node)
	// sendMsg(userId, []byte("欢迎进入聊天系统"))
}

func sendProc(node *Node) {
	for {
		select {
		case data := <-node.DataQueue:
			err := node.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				zap.S().Info("写入消息失败", err)
				return
			}
			fmt.Println("数据发送socket成功")
		}
	}
}

func recProc(node *Node) {
	for {
		//获取信息
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			zap.S().Info("读取消息失败", err)
			return
		}

		// brodMsg(data)

		//这里是简单实现的一种方法
		msg := Message{}
		err = json.Unmarshal(data, &msg)
		if err != nil {
			zap.S().Info("json解析失败", err)
			return
		}

		if msg.Type == 1 {
			zap.S().Info("这是一条私信:", msg.Content)
			tarNode, ok := clientMap[msg.TargetId]
			if !ok {
				zap.S().Info("不存在对应的node", msg.TargetId)
				return
			}

			tarNode.DataQueue <- data
			fmt.Println("发送成功：", string(data))
		}
	}
}
