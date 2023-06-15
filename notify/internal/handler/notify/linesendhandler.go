package notify

import (
	"com.copo/copo_notify/common"
	"com.copo/copo_notify/common/errorz"
	"com.copo/copo_notify/common/response"
	"com.copo/copo_notify/notify/internal/logic/notify"
	"com.copo/copo_notify/notify/internal/svc"
	"com.copo/copo_notify/notify/internal/types"
	"github.com/tal-tech/go-zero/rest/httpx"
	"net/http"
)

func LineSendHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LineSendRequest

		if err := httpx.ParseJsonBody(r, &req); err != nil {
			response.Json(w, r, response.FAIL, nil, err)
			return
		}
		authenticationPaykey := r.Header.Get("authenticationLinekey")
		if isOK, err := common.MicroServiceVerification(authenticationPaykey, ctx.Config.ApiKey.LineKey, ctx.Config.ApiKey.PublicKey); err != nil || !isOK {
			err = errorz.New(response.INTERNAL_SIGN_ERROR)
			response.Json(w, r, err.Error(), nil, err)
			return
		}
		l := notify.NewLineSendLogic(r.Context(), ctx)
		resp, err := l.LineSend(req)
		if err != nil {
			response.Json(w, r, err.Error(), nil, err)
		} else {
			response.Json(w, r, response.SUCCESS, resp, err)
		}
	}
}
