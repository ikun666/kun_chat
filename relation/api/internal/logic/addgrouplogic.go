package logic

import (
	"context"

	"github.com/ikun666/kun_chat/relation/api/internal/svc"
	"github.com/ikun666/kun_chat/relation/api/internal/types"
	"github.com/ikun666/kun_chat/relation/rpc/relationclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddGroupLogic {
	return &AddGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddGroupLogic) AddGroup(req *types.AddGroupRequest) error {
	// todo: add your logic here and delete this line
	_, err := l.svcCtx.RelationRpc.AddGroup(l.ctx, &relationclient.AddGroupRequest{
		Uid:       req.Uid,
		GroupName: req.Name,
	})

	return err
}
