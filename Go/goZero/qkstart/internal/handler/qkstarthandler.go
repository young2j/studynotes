package handler

import (
	"net/http"

	"qkstart/internal/logic"
	"qkstart/internal/svc"
	"qkstart/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func QkstartHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewQkstartLogic(r.Context(), ctx)
		resp, err := l.Qkstart(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
