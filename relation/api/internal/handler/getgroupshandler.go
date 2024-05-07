package handler

import (
	"net/http"

	"github.com/ikun666/kun_chat/relation/api/internal/logic"
	"github.com/ikun666/kun_chat/relation/api/internal/svc"
	"github.com/ikun666/kun_chat/relation/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetGroupsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetGroupsRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetGroupsLogic(r.Context(), svcCtx)
		resp, err := l.GetGroups(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
