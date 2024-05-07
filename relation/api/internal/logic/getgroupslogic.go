package logic

import (
	"context"

	"github.com/ikun666/kun_chat/relation/api/internal/svc"
	"github.com/ikun666/kun_chat/relation/api/internal/types"
	"github.com/ikun666/kun_chat/relation/rpc/pb/relation"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGroupsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetGroupsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGroupsLogic {
	return &GetGroupsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetGroupsLogic) GetGroups(req *types.GetGroupsRequest) (resp *types.GetGroupsResponse, err error) {
	// todo: add your logic here and delete this line
	ret, err := l.svcCtx.RelationRpc.GetGroups(l.ctx, &relation.GetGroupsRequest{
		Uid: req.Uid,
	})
	if err != nil {
		return nil, err
	}
	groups := make([]types.Group, len(ret.Groups))
	for i := range ret.Groups {
		groups[i].Gid = ret.Groups[i].Gid
		groups[i].Name = ret.Groups[i].GroupName
	}
	return &types.GetGroupsResponse{Groups: groups}, nil
}
