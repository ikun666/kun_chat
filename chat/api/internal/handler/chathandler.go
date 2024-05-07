package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
	"github.com/ikun666/kun_chat/chat/api/internal/logic"
	"github.com/ikun666/kun_chat/chat/api/internal/svc"
)

func ChatHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// var req types.ChatRequest
		// if err := httpx.Parse(r, &req); err != nil {
		// 	httpx.ErrorCtx(r.Context(), w, err)
		// 	return
		// }
		// 1.  获取参数 并 检验 token 等合法性
		query := r.URL.Query()
		// fmt.Println("handle:", query)
		Id := query.Get("uid")
		// token := query.Get("token")

		userId, err := strconv.ParseInt(Id, 10, 64)
		if err != nil {
			fmt.Println("类型转换失败", err)
			return
		}
		// 升级为socket
		conn, err := (&websocket.Upgrader{
			//token 校验
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		}).Upgrade(w, r, nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		l := logic.NewChatLogic(r.Context(), svcCtx, conn)
		l.Chat(userId)
		// resp, err := l.Chat(&req)
		// if err != nil {
		// 	httpx.ErrorCtx(r.Context(), w, err)
		// } else {
		// 	httpx.OkJsonCtx(r.Context(), w, resp)
		// }
	}
}
