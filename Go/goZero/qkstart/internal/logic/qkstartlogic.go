package logic

import (
	"context"

	"qkstart/internal/svc"
	"qkstart/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type QkstartLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQkstartLogic(ctx context.Context, svcCtx *svc.ServiceContext) QkstartLogic {
	return QkstartLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QkstartLogic) Qkstart(req types.Request) (*types.Response, error) {
	// todo: add your logic here and delete this line

	return &types.Response{}, nil
}
