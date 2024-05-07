package svc

import (
	"fmt"
	"runtime"

	"github.com/ikun666/kun_chat/chat/api/internal/config"
	"github.com/ikun666/kun_chat/chat/api/internal/global"
	"github.com/ikun666/kun_chat/chat/model"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config        config.Config
	Redis         *redis.Redis
	RelationModel model.RelationModel
	MessagesModel model.MessagesModel
	TaskPool      *global.TaskPool
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DNS)
	ctx := &ServiceContext{
		Config:        c,
		Redis:         redis.MustNewRedis(c.RedisConf),
		RelationModel: model.NewRelationModel(conn),
		MessagesModel: model.NewMessagesModel(conn),
		TaskPool:      global.NewTaskPool(runtime.NumCPU()),
	}
	global.DBOnce.Do(func() {
		global.RelationModel = ctx.RelationModel
		global.MessagesModel = ctx.MessagesModel
		global.Redis = ctx.Redis
		ctx.TaskPool.Run()
		go global.RecordPersistence()
		global.SetSaveKeyFromRedis()
		fmt.Println("init mysql redis")
	})
	return ctx
}
