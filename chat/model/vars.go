package model

import (
	"time"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var ErrNotFound = sqlx.ErrNotFound

type group struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
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
