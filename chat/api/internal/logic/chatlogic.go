package logic

import (
	"context"
	"fmt"

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
		DataQueue: make(chan []byte, 512),
	}

	// 用户关系
	// 将userId和Node绑定
	global.RWLocker.Lock()
	global.ClientMap[userId] = node
	global.RWLocker.Unlock()

	// fmt.Println("uid", userId)

	// 发送消息
	go l.sendProc(node)
	// 接收消息
	go l.recProc(node)
	// sendMsg(userId, []byte("欢迎进入聊天系统"))
	node.DataQueue <- []byte("欢迎进入聊天系统")

}
func (l *ChatLogic) sendProc(node *global.Node) {

	for data := range node.DataQueue {
		err := node.Conn.WriteMessage(websocket.TextMessage, data)
		if err != nil {
			fmt.Println("写入消息失败", err)
			return
		}
	}

}

func (l *ChatLogic) recProc(node *global.Node) {
	for {
		//获取信息
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			fmt.Println("读取消息失败", err)
			return
		}
		l.svcCtx.TaskPool.Add(data)
	}
}
