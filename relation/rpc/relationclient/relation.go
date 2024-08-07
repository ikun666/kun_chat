// Code generated by goctl. DO NOT EDIT.
// Source: relation.proto

package relationclient

import (
	"context"

	"github.com/ikun666/kun_chat/relation/rpc/pb/relation"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddFriendRequest   = relation.AddFriendRequest
	AddFriendResponse  = relation.AddFriendResponse
	AddGroupRequest    = relation.AddGroupRequest
	AddGroupResponse   = relation.AddGroupResponse
	DelFriendRequest   = relation.DelFriendRequest
	DelFriendResponse  = relation.DelFriendResponse
	DelGroupRequest    = relation.DelGroupRequest
	DelGroupResponse   = relation.DelGroupResponse
	GetFriendsRequest  = relation.GetFriendsRequest
	GetFriendsResponse = relation.GetFriendsResponse
	GetGroupsRequest   = relation.GetGroupsRequest
	GetGroupsResponse  = relation.GetGroupsResponse
	Group              = relation.Group

	Relation interface {
		AddFriend(ctx context.Context, in *AddFriendRequest, opts ...grpc.CallOption) (*AddFriendResponse, error)
		DelFriend(ctx context.Context, in *DelFriendRequest, opts ...grpc.CallOption) (*DelFriendResponse, error)
		GetFriends(ctx context.Context, in *GetFriendsRequest, opts ...grpc.CallOption) (*GetFriendsResponse, error)
		AddGroup(ctx context.Context, in *AddGroupRequest, opts ...grpc.CallOption) (*AddGroupResponse, error)
		DelGroup(ctx context.Context, in *DelGroupRequest, opts ...grpc.CallOption) (*DelGroupResponse, error)
		GetGroups(ctx context.Context, in *GetGroupsRequest, opts ...grpc.CallOption) (*GetGroupsResponse, error)
	}

	defaultRelation struct {
		cli zrpc.Client
	}
)

func NewRelation(cli zrpc.Client) Relation {
	return &defaultRelation{
		cli: cli,
	}
}

func (m *defaultRelation) AddFriend(ctx context.Context, in *AddFriendRequest, opts ...grpc.CallOption) (*AddFriendResponse, error) {
	client := relation.NewRelationClient(m.cli.Conn())
	return client.AddFriend(ctx, in, opts...)
}

func (m *defaultRelation) DelFriend(ctx context.Context, in *DelFriendRequest, opts ...grpc.CallOption) (*DelFriendResponse, error) {
	client := relation.NewRelationClient(m.cli.Conn())
	return client.DelFriend(ctx, in, opts...)
}

func (m *defaultRelation) GetFriends(ctx context.Context, in *GetFriendsRequest, opts ...grpc.CallOption) (*GetFriendsResponse, error) {
	client := relation.NewRelationClient(m.cli.Conn())
	return client.GetFriends(ctx, in, opts...)
}

func (m *defaultRelation) AddGroup(ctx context.Context, in *AddGroupRequest, opts ...grpc.CallOption) (*AddGroupResponse, error) {
	client := relation.NewRelationClient(m.cli.Conn())
	return client.AddGroup(ctx, in, opts...)
}

func (m *defaultRelation) DelGroup(ctx context.Context, in *DelGroupRequest, opts ...grpc.CallOption) (*DelGroupResponse, error) {
	client := relation.NewRelationClient(m.cli.Conn())
	return client.DelGroup(ctx, in, opts...)
}

func (m *defaultRelation) GetGroups(ctx context.Context, in *GetGroupsRequest, opts ...grpc.CallOption) (*GetGroupsResponse, error) {
	client := relation.NewRelationClient(m.cli.Conn())
	return client.GetGroups(ctx, in, opts...)
}
