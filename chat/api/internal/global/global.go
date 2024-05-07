package global

import (
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/ikun666/kun_chat/chat/model"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type Message struct {
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
	Conn *websocket.Conn //连接
	// Addr      string          //客户端地址
	DataQueue chan []byte //消息
}

// 映射关系
var ClientMap = make(map[int64]*Node)

// 读写锁
var RWLocker sync.RWMutex

// var UDPChan chan []byte = make(chan []byte, 1024)
var RelationModel model.RelationModel
var MessagesModel model.MessagesModel
var Redis *redis.Redis
var DBOnce sync.Once

var saveKeyMutex sync.RWMutex

// 记录用户聊天key set
var saveKey = make(map[string]struct{})

// 群id和好友id界限
var groupLimit int64 = 1000000000

// 过期时间为10天
var expirationTime = 7 * 24 * 60 * 60

// 启动定时任务，每隔24小时执行一次
var interval time.Duration = 24
