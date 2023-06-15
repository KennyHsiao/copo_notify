package notify

import (
	"context"

	"com.copo/copo_notify/notify/internal/svc"
	"github.com/tal-tech/go-zero/core/logx"
)

type LineAuthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLineAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) LineAuthLogic {
	return LineAuthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LineAuthLogic) LineAuth() error {
	// todo: add your logic here and delete this line

	return nil
}
