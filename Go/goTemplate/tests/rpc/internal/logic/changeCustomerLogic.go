package logic

import (
	"context"
	"github.com/globalsign/mgo/bson"
	"github.com/young2j/gocopy"
	"net/http"
	"scana/common/utils"

	"go-template/tests/rpc/internal/svc"
	"go-template/tests/rpc/userspb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangeCustomerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewChangeCustomerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangeCustomerLogic {
	return &ChangeCustomerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ChangeCustomerLogic) ChangeCustomer(in *userspb.ChangeCustomerReq) (*userspb.ChangeCustomerResp, error) {
	data := bson.M{}
	gocopy.CopyWithOption(&data, in, &gocopy.Option{
		IgnoreZero: true,
		NameFromTo: map[string]string{"Id": "_id"},
		Converters: utils.CONVERTER_TO_COMPLEX,
	})
	err := l.svcCtx.UsersCustomerModel.UpdateOneId(l.ctx, in.Id, data)
	if err != nil {
		return nil, err
	}
	return &userspb.ChangeCustomerResp{
		Code: http.StatusOK,
		Msg:  "修改成功",
	}, nil

}
