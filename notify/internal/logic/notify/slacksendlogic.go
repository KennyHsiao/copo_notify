package notify

import (
	"context"
	"fmt"
	"github.com/slack-go/slack"

	"com.copo/copo_notify/notify/internal/svc"
	"com.copo/copo_notify/notify/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type SlackSendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSlackSendLogic(ctx context.Context, svcCtx *svc.ServiceContext) SlackSendLogic {
	return SlackSendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SlackSendLogic) SlackSend(req types.SlackSendRequest) (*types.SlackSendResponse, error) {
	api := slack.New(l.svcCtx.Config.Notify.Slack.BotToken)

	message, s, err := api.PostMessage(req.ChatID, slack.MsgOptionText(req.Message, false))
	if err != nil {
		return nil, err
	}

	return &types.SlackSendResponse{
		Msg: fmt.Sprintf("%s at %s", message, s),
	}, nil
}
