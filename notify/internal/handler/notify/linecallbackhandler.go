package notify

import (
	"com.copo/copo_notify/notify/internal/svc"
	"context"
	"fmt"
	"github.com/utahta/go-linenotify/auth"
	"github.com/utahta/go-linenotify/token"
	"net/http"
)

func LineCallbackHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := auth.ParseRequest(r)
		if err != nil {
			fmt.Fprintf(w, "error:%v", err)
			return
		}

		state, err := r.Cookie("state")
		if err != nil {
			fmt.Fprintf(w, "error:%v", err)
			return
		}

		if resp.State != state.Value {
			fmt.Fprintf(w, "error:%v", err)
			return
		}

		c := token.NewClient(
			ctx.Config.Notify.Line.CallbackURL,
			ctx.Config.Notify.Line.ClientID,
			ctx.Config.Notify.Line.ClientSecret,
		)
		accessToken, err := c.GetAccessToken(context.Background(), resp.Code)
		if err != nil {
			fmt.Fprintf(w, "error:%v", err)
			return
		}

		fmt.Fprintf(w, "token:%v", accessToken)
	}
}
