package db

import (
	"github.com/ikun666/kun_chat/chat/model"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var DNS = "root:123456@tcp(192.168.44.132:3306)/kun_chat?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai"
var conn = sqlx.NewMysql(DNS)

// var MessagesModel = model.NewMessagesModel(conn)
// var GroupModel = model.NewGroupModel(conn)
var RelationModel = model.NewRelationModel(conn)

var opt = redis.Options{
	Addr:     "192.168.44.132:6379", // redis地址
	Password: "",                    // redis密码，没有则留空
	DB:       10,                    // 默认数据库，默认是0
}
var RedisDB = redis.NewClient(&opt)
