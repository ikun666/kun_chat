package logic

import (
	"context"

	"github.com/ikun666/kun_chat/relation/rpc/internal/svc"
	"github.com/ikun666/kun_chat/relation/rpc/pb/relation"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelGroupLogic {
	return &DelGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelGroupLogic) DelGroup(in *relation.DelGroupRequest) (*relation.DelGroupResponse, error) {
	// todo: add your logic here and delete this line
	err := l.svcCtx.GroupModel.DeleteByUidName(l.ctx, in.Uid, in.GroupName)
	if err != nil {
		return nil, err
	}
	return &relation.DelGroupResponse{}, nil
}
