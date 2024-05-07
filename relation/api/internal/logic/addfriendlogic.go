package logic

import (
	"context"

	"github.com/ikun666/kun_chat/relation/api/internal/svc"
	"github.com/ikun666/kun_chat/relation/api/internal/types"
	"github.com/ikun666/kun_chat/relation/rpc/relationclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddFriendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddFriendLogic {
	return &AddFriendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddFriendLogic) AddFriend(req *types.AddFriendRequest) error {
	// todo: add your logic here and delete this line
	_, err := l.svcCtx.RelationRpc.AddFriend(l.ctx, &relationclient.AddFriendRequest{
		Uid:  req.Uid,
		Tid:  req.Tid,
		Type: req.Type,
	})
	return err
}
