package notify

import (
	"com.copo/copo_notify/common/response"
	"com.copo/copo_notify/notify/internal/logic/notify"
	"com.copo/copo_notify/notify/internal/svc"
	"com.copo/copo_notify/notify/internal/types"
	"github.com/tal-tech/go-zero/rest/httpx"
	"net/http"
)

func TelegramSendHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TelegramSendRequest

		if err := httpx.ParseJsonBody(r, &req); err != nil {
			response.Json(w, r, response.FAIL, nil, err)
			return
		}

		l := notify.NewTelegramSendLogic(r.Context(), ctx)
		resp, err := l.TelegramSend(req)
		if err != nil {
			response.Json(w, r, err.Error(), nil, err)
		} else {
			response.Json(w, r, response.SUCCESS, resp, err)
		}
	}
}
