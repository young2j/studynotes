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

type UpsertCustomerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpsertCustomerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpsertCustomerLogic {
	return &UpsertCustomerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpsertCustomerLogic) UpsertCustomer(in *userspb.UpsertCustomerReq) (*userspb.UpsertCustomerResp, error) {
	query := bson.M{}
	gocopy.CopyWithOption(&query, in.Query, &gocopy.Option{
		IgnoreZero: true,
		NameFromTo: map[string]string{"Id": "_id"},
		Converters: utils.CONVERTER_TO_COMPLEX,
	})

	changeInfo, err := l.svcCtx.UsersCustomerModel.Upsert(l.ctx, query, in.Data)
	if err != nil {
		return nil, err
	}
	data := ""
	objectId, ok := changeInfo.UpsertedId.(bson.ObjectId)
	if ok {
		data = objectId.Hex()
	}

	return &userspb.UpsertCustomerResp{
		Code: http.StatusOK,
		Msg:  "操作成功",
		Data: data,
	}, nil

}
