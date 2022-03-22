package logic

import (
	"context"
	"github.com/young2j/gocopy"
	"net/http"
	"go-template/tests/model"

	"go-template/tests/rpc/internal/svc"
	"go-template/tests/rpc/userspb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCustomerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCustomerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCustomerLogic {
	return &AddCustomerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddCustomerLogic) AddCustomer(in *userspb.AddCustomerReq) (*userspb.AddCustomerResp, error) {
	data := &model.UsersCustomer{}
	gocopy.Copy(data, in)
	err := l.svcCtx.UsersCustomerModel.Insert(l.ctx, data)
	if err != nil {
		return nil, err
	}

	return &userspb.AddCustomerResp{
		Code: http.StatusOK,
		Msg:  "添加成功",
	}, nil

}
