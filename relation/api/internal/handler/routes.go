// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"github.com/ikun666/kun_chat/relation/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/relation/addFriend",
				Handler: AddFriendHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/relation/addGroup",
				Handler: AddGroupHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/relation/delFriend",
				Handler: DelFriendHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/relation/delGroup",
				Handler: DelGroupHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/relation/getFriends",
				Handler: GetFriendsHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/relation/getGroups",
				Handler: GetGroupsHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}
