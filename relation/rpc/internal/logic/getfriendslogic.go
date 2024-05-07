package logic

import (
	"context"

	"github.com/ikun666/kun_chat/relation/rpc/internal/svc"
	"github.com/ikun666/kun_chat/relation/rpc/pb/relation"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFriendsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFriendsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFriendsLogic {
	return &GetFriendsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFriendsLogic) GetFriends(in *relation.GetFriendsRequest) (*relation.GetFriendsResponse, error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.RelationModel.GetFriends(l.ctx, in.Uid, in.Type)
	if err != nil {
		return nil, err
	}
	return &relation.GetFriendsResponse{Tid: res}, nil
}
