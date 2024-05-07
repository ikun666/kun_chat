package logic

import (
	"context"

	"github.com/ikun666/kun_chat/relation/api/internal/svc"
	"github.com/ikun666/kun_chat/relation/api/internal/types"
	"github.com/ikun666/kun_chat/relation/rpc/relationclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelFriendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelFriendLogic {
	return &DelFriendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelFriendLogic) DelFriend(req *types.DelFriendRequest) error {
	// todo: add your logic here and delete this line
	_, err := l.svcCtx.RelationRpc.DelFriend(l.ctx, &relationclient.DelFriendRequest{
		Uid:  req.Uid,
		Tid:  req.Tid,
		Type: req.Type,
	})
	return err
}
