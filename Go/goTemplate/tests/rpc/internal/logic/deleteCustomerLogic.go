package logic

import (
	"context"
	"net/http"

	"go-template/tests/rpc/internal/svc"
	"go-template/tests/rpc/userspb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCustomerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCustomerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCustomerLogic {
	return &DeleteCustomerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteCustomerLogic) DeleteCustomer(in *userspb.DeleteCustomerReq) (*userspb.DeleteCustomerResp, error) {
	err := l.svcCtx.UsersCustomerModel.RemoveId(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &userspb.DeleteCustomerResp{
		Code: http.StatusOK,
		Msg:  "删除成功",
	}, nil

}
