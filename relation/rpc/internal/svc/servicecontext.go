package svc

import (
	"github.com/ikun666/kun_chat/relation/model"
	"github.com/ikun666/kun_chat/relation/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config        config.Config
	RelationModel model.RelationModel
	GroupModel    model.GroupModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		RelationModel: model.NewRelationModel(sqlx.NewMysql(c.DNS)),
		GroupModel:    model.NewGroupModel(sqlx.NewMysql(c.DNS)),
	}
}
