package svc

import (
	"github.com/ikun666/kun_chat/relation/api/internal/config"
	"github.com/ikun666/kun_chat/relation/rpc/relationclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	RelationRpc relationclient.Relation
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		RelationRpc: relationclient.NewRelation(zrpc.MustNewClient(c.RelationRpc)),
	}
}
