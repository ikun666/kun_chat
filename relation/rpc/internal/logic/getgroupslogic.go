package logic

import (
	"context"

	"github.com/ikun666/kun_chat/relation/rpc/internal/svc"
	"github.com/ikun666/kun_chat/relation/rpc/pb/relation"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGroupsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGroupsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGroupsLogic {
	return &GetGroupsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetGroupsLogic) GetGroups(in *relation.GetGroupsRequest) (*relation.GetGroupsResponse, error) {
	// todo: add your logic here and delete this line
	resp, err := l.svcCtx.GroupModel.GetGroups(l.ctx, in.Uid)
	if err != nil {
		return nil, err
	}
	groups := make([]*relation.Group, len(resp))
	for i := range resp {
		groups[i] = &relation.Group{
			Gid:       resp[i].Id,
			GroupName: resp[i].Name,
		}
	}
	return &relation.GetGroupsResponse{Groups: groups}, nil
}
