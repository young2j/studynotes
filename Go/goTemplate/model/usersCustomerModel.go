package model

import (
	"context"
	"strings"
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	cachec "github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/mongoc"
)

var prefixUsersCustomerCacheKey = "cache:UsersCustomer:"

type Operator struct {
	Id   bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name string        `json:"name" bson:"name" description:"姓名"`
}

// 客户模型
type UsersCustomer struct {
	Id             bson.ObjectId //
	Uid            string        // uuid
	Identity       int32         // 客户身份
	Name           string        // 姓名
	Status         int32         // 客户状态
	AccountName    string        // 账号名称
	AccountStatus  int32         // 账号状态
	Email          string        // 客户邮箱
	Phone          string        // 客户电话
	KsAccount      string        // 创宇通行证
	SalesPerson    string        // 关联销售人员
	CanCustomizeKW int32         // 是否可自定义敏感词
	CreatedTime    time.Time     // 创建时间
	UpdatedTime    time.Time     // 更新时间
	Operator       Operator      // 操作人，{id:'',name:''}
	Domains        []string      // 监测域名
}

// type UsersCustomer struct {
// 	Id             bson.ObjectId `json:"id" bson:"_id,omitempty"`
// 	Uid            string        `json:"uid" bson:"uid" description:"uuid"`
// 	Identity       int32         `json:"identity" bson:"identity" description:"客户身份"`
// 	Name           string        `json:"name" bson:"name" description:"姓名"`
// 	Status         int32         `json:"status" bson:"status" description:"客户状态"`
// 	AccountName    string        `json:"accountName" bson:"accountName" description:"账号名称"`
// 	AccountStatus  int32         `json:"accountStatus" bson:"accountStatus" description:"账号状态"`
// 	Email          string        `json:"email" bson:"email" description:"客户邮箱"`
// 	Phone          string        `json:"phone" bson:"phone" description:"客户电话"`
// 	KsAccount      string        `json:"ksAccount" bson:"ksAccount" description:"创宇通行证"`
// 	SalesPerson    string        `json:"salesPerson" bson:"salesPerson" description:"关联销售人员"`
// 	CanCustomizeKW int32         `json:"canCustomizeKW" bson:"canCustomizeKW" description:"是否可自定义敏感词"`
// 	CreatedTime    time.Time     `json:"createdTime" bson:"createdTime" description:"创建时间"`
// 	UpdatedTime    time.Time     `json:"updatedTime" bson:"updatedTime" description:"更新时间"`
// 	Operator       Operator      `json:"operator" bson:"operator" description:"操作人，{id:'',name:''}"`
// 	Domains        []string      `json:"domains" bson:"domains" description:"监测域名"`
// }

type UsersCustomerModel interface {
	cacheKeyFromQuery(query interface{}) (string, error)
	Count(ctx context.Context, query interface{}) (int32, error)
	Insert(ctx context.Context, data *UsersCustomer) error
	FindOne(ctx context.Context, query interface{}) (*UsersCustomer, error)
	FindOneId(ctx context.Context, id string) (*UsersCustomer, error)
	FindAll(ctx context.Context, query interface{}, opts ...mongoc.QueryOption) ([]*UsersCustomer, error)
	UpdateOne(ctx context.Context, data *UsersCustomer) error
	Update(ctx context.Context, query interface{}, data *UsersCustomer) error
	Upsert(ctx context.Context, query interface{}, data *UsersCustomer) (*mgo.ChangeInfo, error)
	RemoveId(ctx context.Context, id string) error
	RemoveAll(ctx context.Context, query interface{}) (*mgo.ChangeInfo, error)
	Remove(ctx context.Context, query interface{}) error
}

type defaultUsersCustomerModel struct {
	*mongoc.Model
}

func NewUsersCustomerModel(url, collection string, c cachec.CacheConf) UsersCustomerModel {
	return &defaultUsersCustomerModel{
		Model: mongoc.MustNewModel(url, collection, c),
	}
}

func (m *defaultUsersCustomerModel) cacheKeyFromQuery(query interface{}) (string, error) {
	mj, err := bson.MarshalJSON(query)
	if err != nil {
		return "", err
	}

	b := strings.Builder{}
	_, err = b.WriteString(prefixUsersCustomerCacheKey)
	if err != nil {
		return "", err
	}

	_, err = b.Write(mj)
	if err != nil {
		return "", err
	}

	return b.String(), nil
}

func (m *defaultUsersCustomerModel) Count(ctx context.Context, query interface{}) (int32, error) {
	session, err := m.TakeSession()
	if err != nil {
		return 0, err
	}
	defer m.PutSession(session)

	count, err := m.GetCollection(session).Count(query)
	if err != nil {
		return 0, err
	}
	return int32(count), nil
}

func (m *defaultUsersCustomerModel) Insert(ctx context.Context, data *UsersCustomer) error {
	if !data.Id.Valid() {
		data.Id = bson.NewObjectId()
	}

	session, err := m.TakeSession()
	if err != nil {
		return err
	}
	defer m.PutSession(session)

	return m.GetCollection(session).Insert(data)
}

func (m *defaultUsersCustomerModel) FindOne(ctx context.Context, query interface{}) (*UsersCustomer, error) {
	session, err := m.TakeSession()
	if err != nil {
		return nil, err
	}
	defer m.PutSession(session)

	var data UsersCustomer

	key, err := m.cacheKeyFromQuery(query)
	if err != nil {
		return nil, err
	}
	err = m.GetCollection(session).FindOne(&data, key, query)

	switch err {
	case nil:
		return &data, nil
	case mongoc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUsersCustomerModel) FindOneId(ctx context.Context, id string) (*UsersCustomer, error) {
	if !bson.IsObjectIdHex(id) {
		return nil, ErrInvalidObjectId
	}

	session, err := m.TakeSession()
	if err != nil {
		return nil, err
	}
	defer m.PutSession(session)

	var data UsersCustomer

	key := prefixUsersCustomerCacheKey + id
	err = m.GetCollection(session).FindOneId(&data, key, bson.ObjectIdHex(id))
	switch err {
	case nil:
		return &data, nil
	case mongoc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUsersCustomerModel) FindAll(ctx context.Context, query interface{}, opts ...mongoc.QueryOption) ([]*UsersCustomer, error) {

	session, err := m.TakeSession()
	if err != nil {
		return nil, err
	}

	defer m.PutSession(session)
	var data []*UsersCustomer

	err = m.GetCollection(session).FindAllNoCache(&data, query, opts...)
	switch err {
	case nil:
		return data, nil
	case mongoc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUsersCustomerModel) UpdateOne(ctx context.Context, data *UsersCustomer) error {
	session, err := m.TakeSession()
	if err != nil {
		return err
	}
	defer m.PutSession(session)

	key := prefixUsersCustomerCacheKey + data.Id.Hex()
	err = m.GetCollection(session).UpdateId(data.Id, data, key)

	return err
}

func (m *defaultUsersCustomerModel) Update(ctx context.Context, query interface{}, data *UsersCustomer) error {
	session, err := m.TakeSession()
	if err != nil {
		return err
	}
	defer m.PutSession(session)

	key, err := m.cacheKeyFromQuery(query)
	if err != nil {
		return err
	}
	m.GetCollection(session).Update(query, data, key)
	return err
}

func (m *defaultUsersCustomerModel) Upsert(ctx context.Context, query interface{}, data *UsersCustomer) (*mgo.ChangeInfo, error) {
	session, err := m.TakeSession()
	if err != nil {
		return nil, err
	}
	defer m.PutSession(session)

	key, err := m.cacheKeyFromQuery(query)
	if err != nil {
		return nil, err
	}
	upsertData := bson.M{"$set": data}
	changeInfo, err := m.GetCollection(session).Upsert(query, upsertData, key)
	return changeInfo, err
}

func (m *defaultUsersCustomerModel) RemoveId(ctx context.Context, id string) error {
	if !bson.IsObjectIdHex(id) {
		return ErrInvalidObjectId
	}

	session, err := m.TakeSession()
	if err != nil {
		return err
	}
	defer m.PutSession(session)

	key := prefixUsersCustomerCacheKey + id
	err = m.GetCollection(session).RemoveId(bson.ObjectIdHex(id), key)
	return err
}

func (m *defaultUsersCustomerModel) RemoveAll(ctx context.Context, query interface{}) (*mgo.ChangeInfo, error) {
	session, err := m.TakeSession()
	if err != nil {
		return nil, err
	}
	defer m.PutSession(session)

	key, err := m.cacheKeyFromQuery(query)
	if err != nil {
		return nil, err
	}
	changeInfo, err := m.GetCollection(session).RemoveAll(query, key)
	return changeInfo, err
}

func (m *defaultUsersCustomerModel) Remove(ctx context.Context, query interface{}) error {
	session, err := m.TakeSession()
	if err != nil {
		return err
	}
	defer m.PutSession(session)

	key, err := m.cacheKeyFromQuery(query)
	if err != nil {
		return err
	}
	err = m.GetCollection(session).Remove(query, key)

	return err
}
