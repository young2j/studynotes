package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type KeywordsKeywords struct {
	Id           int              `orm:"column(id);auto"`
	KwCategory   string           `orm:"column(kw_category);size(20);null"`
	Word         string           `orm:"column(word);size(100)"`
	OperatorName string           `orm:"column(operator_name);size(20);null"`
	OperateTime  time.Time        `orm:"column(operate_time);type(datetime);null"`
	Remark       string           `orm:"column(remark);size(200);null"`
	RiskLevel    int              `orm:"column(risk_level)"`
	Effective    string           `orm:"column(effective);size(20)"`
	KwLabel2     string           `orm:"column(kw_label2);size(20)"`
	Status       int              `orm:"column(status)"`
	ResultTime   time.Time        `orm:"column(result_time);type(datetime);null"`
	StatusAudit  string           `orm:"column(status_audit);size(10)"`
	UserId       *MainUserprofile `orm:"column(user_id);rel(fk)"`
	ReceiveTime  time.Time        `orm:"column(receive_time);type(datetime);null"`
	CreateTime   time.Time        `orm:"column(create_time);type(datetime);null"`
}

func (t *KeywordsKeywords) TableName() string {
	return "keywords_keywords"
}

func init() {
	orm.RegisterModel(new(KeywordsKeywords))
}

// AddKeywordsKeywords insert a new KeywordsKeywords into database and returns
// last inserted Id on success.
func AddKeywordsKeywords(m *KeywordsKeywords) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetKeywordsKeywordsById retrieves KeywordsKeywords by Id. Returns error if
// Id doesn't exist
func GetKeywordsKeywordsById(id int) (v *KeywordsKeywords, err error) {
	o := orm.NewOrm()
	v = &KeywordsKeywords{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllKeywordsKeywords retrieves all KeywordsKeywords matches certain condition. Returns empty list if
// no records exist
func GetAllKeywordsKeywords(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(KeywordsKeywords))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []KeywordsKeywords
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateKeywordsKeywords updates KeywordsKeywords by Id and returns error if
// the record to be updated doesn't exist
func UpdateKeywordsKeywordsById(m *KeywordsKeywords) (err error) {
	o := orm.NewOrm()
	v := KeywordsKeywords{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteKeywordsKeywords deletes KeywordsKeywords by Id and returns error if
// the record to be deleted doesn't exist
func DeleteKeywordsKeywords(id int) (err error) {
	o := orm.NewOrm()
	v := KeywordsKeywords{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&KeywordsKeywords{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
