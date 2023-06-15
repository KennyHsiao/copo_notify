package notify

import (
	"com.copo/copo_notify/notify/internal/svc"
	"com.copo/copo_notify/notify/internal/types"
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"

	"github.com/zeromicro/go-zero/core/logx"
)

type TelegramNotifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTelegramNotifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TelegramNotifyLogic {
	return &TelegramNotifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TelegramNotifyLogic) TelegramNotify(req *types.TelegramNotifyRequest) error {
	notifyChatId := l.svcCtx.Config.Notify.Telegram.NotifyChatId

	bot, err := tgbotapi.NewBotAPI(l.svcCtx.Config.Notify.Telegram.BotToken)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	logx.WithContext(l.ctx).Info("发送群组id: ", notifyChatId)
	msg := tgbotapi.NewMessage(int64(notifyChatId), req.Message)

	_, errs := bot.Send(msg)

	if errs != nil {
		return errs
	}

	return nil
}
