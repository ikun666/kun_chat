package logic

import (
	"context"

	"github.com/ikun666/kun_chat/relation/api/internal/svc"
	"github.com/ikun666/kun_chat/relation/api/internal/types"
	"github.com/ikun666/kun_chat/relation/rpc/pb/relation"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelGroupLogic {
	return &DelGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelGroupLogic) DelGroup(req *types.DelGroupRequest) error {
	// todo: add your logic here and delete this line
	_, err := l.svcCtx.RelationRpc.DelGroup(l.ctx, &relation.DelGroupRequest{
		Uid:       req.Uid,
		GroupName: req.Name,
	})
	return err
}
