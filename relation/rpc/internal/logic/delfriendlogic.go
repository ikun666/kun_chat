package logic

import (
	"context"

	"github.com/ikun666/kun_chat/relation/rpc/internal/svc"
	"github.com/ikun666/kun_chat/relation/rpc/pb/relation"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelFriendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelFriendLogic {
	return &DelFriendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelFriendLogic) DelFriend(in *relation.DelFriendRequest) (*relation.DelFriendResponse, error) {
	// todo: add your logic here and delete this line
	err := l.svcCtx.RelationModel.DelFriend(l.ctx, in.Uid, in.Tid, in.Type)
	if err != nil {
		return nil, err
	}

	return &relation.DelFriendResponse{}, nil
}
