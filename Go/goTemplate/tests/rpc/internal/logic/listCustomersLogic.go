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

type ListCustomersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListCustomersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListCustomersLogic {
	return &ListCustomersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListCustomersLogic) ListCustomers(in *userspb.ListCustomersReq) (*userspb.ListCustomersResp, error) {
	opts := utils.GetQueryOption(&utils.QueryOption{
		Page:          in.Page,
		PageSize:      in.PageSize,
		SortKeys:      in.SortKeys,
		ProjectFields: in.ProjectFields,
		ExcludeFields: in.ExcludeFields,
	})
	q := bson.M{}
	if in.Search != "" {
		q["$or"] = []bson.M{
			{"field": bson.M{"$regex": in.Search}},
		}
	}
	count, err := l.svcCtx.UsersCustomerModel.Count(l.ctx, q)
	if err != nil {
		return nil, err
	}
	customer, err := l.svcCtx.UsersCustomerModel.FindAll(l.ctx, q, opts)
	if err != nil {
		return nil, err
	}

	data := []*userspb.CustomerInfo{}
	gocopy.CopyWithOption(&data, customer, &gocopy.Option{
		NameFromTo: map[string]string{"_id": "Id"},
		Converters: utils.CONVERTER_TO_STRING,
	})

	return &userspb.ListCustomersResp{
		Code:      http.StatusOK,
		Msg:       "查询成功",
		Data:      data,
		Count:     count,
		TotalPage: utils.GetTotalPage(count, in.PageSize),
	}, nil

}
