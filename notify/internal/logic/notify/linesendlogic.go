package notify

import (
	"context"

	"com.copo/copo_notify/notify/internal/svc"
	"com.copo/copo_notify/notify/internal/types"

	"github.com/tal-tech/go-zero/core/logx"

	"github.com/utahta/go-linenotify"
)

type LineSendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLineSendLogic(ctx context.Context, svcCtx *svc.ServiceContext) LineSendLogic {
	return LineSendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LineSendLogic) LineSend(req types.LineSendRequest) (*types.LineSendResponse, error) {
	token := "pYjUPCj9eUZRN6ULdEvgoc0xb7cV8ys5xX90qaY8GoA"
	c := linenotify.NewClient()
	notify, err := c.Notify(context.Background(), token, req.Message, "", "", nil)
	if err != nil {
		return nil, err
	}
	return &types.LineSendResponse{
		Msg: notify.Message,
	}, nil
}
