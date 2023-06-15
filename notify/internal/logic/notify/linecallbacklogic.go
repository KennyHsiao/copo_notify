package notify

import (
	"context"

	"com.copo/copo_notify/notify/internal/svc"
	"github.com/tal-tech/go-zero/core/logx"
)

type LineCallbackLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLineCallbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) LineCallbackLogic {
	return LineCallbackLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LineCallbackLogic) LineCallback() error {
	// todo: add your logic here and delete this line

	return nil
}
