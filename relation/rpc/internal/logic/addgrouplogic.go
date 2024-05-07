package logic

import (
	"context"

	"github.com/ikun666/kun_chat/relation/model"
	"github.com/ikun666/kun_chat/relation/rpc/internal/svc"
	"github.com/ikun666/kun_chat/relation/rpc/pb/relation"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddGroupLogic {
	return &AddGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddGroupLogic) AddGroup(in *relation.AddGroupRequest) (*relation.AddGroupResponse, error) {
	// todo: add your logic here and delete this line
	_, err := l.svcCtx.GroupModel.FindOneByName(l.ctx, in.GroupName)
	if err != nil && err != model.ErrNotFound {
		return nil, err
	}
	if err != model.ErrNotFound {
		return nil, status.Error(1234, "已有该群")
	}
	_, err = l.svcCtx.GroupModel.Insert(l.ctx, &model.Group{
		OwnerId: in.Uid,
		Name:    in.GroupName,
	})
	if err != nil {
		return nil, err
	}
	return &relation.AddGroupResponse{}, nil
}
