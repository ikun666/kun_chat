package handler

import (
	"net/http"

	"github.com/ikun666/kun_chat/chat/api/internal/logic"
	"github.com/ikun666/kun_chat/chat/api/internal/svc"
	"github.com/ikun666/kun_chat/chat/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetMsgsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetMsgRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetMsgsLogic(r.Context(), svcCtx)
		resp, err := l.GetMsgs(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
