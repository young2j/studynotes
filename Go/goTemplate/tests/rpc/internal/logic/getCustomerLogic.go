package logic

import (
	"context"
	"github.com/globalsign/mgo/bson"
	"github.com/young2j/gocopy"
	"net/http"
	"scana/common/utils"
	"go-template/tests/model"

	"go-template/tests/rpc/internal/svc"
	"go-template/tests/rpc/userspb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCustomerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCustomerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCustomerLogic {
	return &GetCustomerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCustomerLogic) GetCustomer(in *userspb.GetCustomerReq) (*userspb.GetCustomerResp, error) {
	var (
		customer *model.UsersCustomer
		err      error
	)
	if in.Id != "" {
		customer, err = l.svcCtx.UsersCustomerModel.FindOneId(l.ctx, in.Id)
		if err != nil {
			return nil, err
		}
	} else {
		q := bson.M{}
		customer, err = l.svcCtx.UsersCustomerModel.FindOne(l.ctx, q)
		if err != nil {
			return nil, err
		}
	}

	data := &userspb.CustomerInfo{}
	gocopy.CopyWithOption(data, customer, &gocopy.Option{
		NameFromTo: map[string]string{"_id": "Id"},
		Converters: utils.CONVERTER_TO_STRING,
	})

	return &userspb.GetCustomerResp{
		Code: http.StatusOK,
		Msg:  "查询成功",
		Data: data,
	}, nil

}
