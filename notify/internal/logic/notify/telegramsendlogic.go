package notify

import (
	"com.copo/copo_notify/notify/internal/svc"
	"com.copo/copo_notify/notify/internal/types"
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/tal-tech/go-zero/core/logx"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type TelegramSendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTelegramSendLogic(ctx context.Context, svcCtx *svc.ServiceContext) TelegramSendLogic {
	return TelegramSendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TelegramSendLogic) TelegramSend(req types.TelegramSendRequest) (*types.TelegramSendResponse, error) {
	bot, err := tgbotapi.NewBotAPI(l.svcCtx.Config.Notify.Telegram.BotToken)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	chatId, _ := strconv.Atoi(req.ChatID)
	msg := tgbotapi.NewMessage(int64(chatId), req.Message)

	message, errs := bot.Send(msg)

	if errs != nil {
		return nil, errs
	}

	//editMsg := tgbotapi.NewEditMessageText(int64(chatId), 11, "测试测试222")
	//
	//backMsgE, errE := bot.Send(editMsg)
	//
	//if errE != nil {
	//	return nil, errE
	//}

	//delMsg := tgbotapi.NewDeleteMessage(int64(chatId), 10)
	//_, errD := bot.DeleteMessage(delMsg)
	//if errD != nil {
	//	return nil, errD
	//}

	//a:=tgbotapi.NewMessageToChannel("TEST", "hk4g4")
	//
	//_, errC := bot.Send(a)
	//if errC != nil {
	//	return nil, errC
	//}
	//
	//u:=tgbotapi.NewUpdate(0)
	//updates , errU := bot.GetUpdatesChan(u)
	//if errU != nil {
	//	return nil, errU
	//}

	return &types.TelegramSendResponse{
		//Msg: fmt.Sprintf("Authorized on account %s", bot.Self.UserName),
		Msg: fmt.Sprintf("%d", message.MessageID),
	}, nil
}
