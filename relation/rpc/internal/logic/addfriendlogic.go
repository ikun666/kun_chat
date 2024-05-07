package logic

import (
	"context"
	"errors"

	"github.com/ikun666/kun_chat/relation/model"
	"github.com/ikun666/kun_chat/relation/rpc/internal/svc"
	"github.com/ikun666/kun_chat/relation/rpc/pb/relation"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddFriendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddFriendLogic {
	return &AddFriendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddFriendLogic) AddFriend(in *relation.AddFriendRequest) (*relation.AddFriendResponse, error) {
	// todo: add your logic here and delete this line
	ok, err := l.svcCtx.RelationModel.FindFriendByUidTidType(l.ctx, in.Uid, in.Tid, in.Type)
	if err != nil && err != model.ErrNotFound {
		return nil, err
	}
	if ok {
		if in.Type == 1 {
			return nil, errors.New("已经是好友")
		} else if in.Type == 2 {
			return nil, errors.New("已经在群里")
		} else {
			return nil, errors.New("type err")
		}

	}
	err = l.svcCtx.RelationModel.AddFriend(l.ctx, in.Uid, in.Tid, in.Type)
	if err != nil {
		return nil, err
	}

	return &relation.AddFriendResponse{}, nil
}
