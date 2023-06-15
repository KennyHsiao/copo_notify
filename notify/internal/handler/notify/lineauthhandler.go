package notify

import (
	"com.copo/copo_notify/notify/internal/svc"
	"fmt"
	"github.com/utahta/go-linenotify/auth"
	"net/http"
	"time"
)

func LineAuthHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c, err := auth.New(
			ctx.Config.Notify.Line.ClientID,
			ctx.Config.Notify.Line.CallbackURL,
		)

		if err != nil {
			fmt.Fprintf(w, "error:%v", err)
			return
		}

		http.SetCookie(w, &http.Cookie{Name: "state", Value: c.State, Expires: time.Now().Add(60 * time.Second)})

		c.Redirect(w, r)

	}
}
