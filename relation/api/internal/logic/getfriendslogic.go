package logic

import (
	"context"

	"github.com/ikun666/kun_chat/relation/api/internal/svc"
	"github.com/ikun666/kun_chat/relation/api/internal/types"
	"github.com/ikun666/kun_chat/relation/rpc/pb/relation"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFriendsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFriendsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFriendsLogic {
	return &GetFriendsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFriendsLogic) GetFriends(req *types.GetFriendsRequest) (resp *types.GetFriendsResponse, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.RelationRpc.GetFriends(l.ctx, &relation.GetFriendsRequest{
		Uid:  req.Uid,
		Type: req.Type,
	})
	if err != nil {
		return nil, err
	}
	return &types.GetFriendsResponse{Tid: res.Tid}, nil
}
