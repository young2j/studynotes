package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type MainUserprofile struct {
	Id              int                 `orm:"column(id);auto"`
	Name            string              `orm:"column(name);size(30)"`
	Avatar          string              `orm:"column(avatar);size(255)"`
	EmployeeId      string              `orm:"column(employee_id);size(20)"`
	Base            string              `orm:"column(base);size(30)"`
	Phone           string              `orm:"column(phone);size(30)"`
	Status          string              `orm:"column(status);size(50)"`
	StaffType       string              `orm:"column(staff_type);size(50)"`
	Level           string              `orm:"column(level);size(50)"`
	CreateTime      time.Time           `orm:"column(create_time);type(datetime)"`
	UpdateTime      time.Time           `orm:"column(update_time);type(datetime)"`
	FaceUpload      int                 `orm:"column(face_upload)"`
	MaxSkipNum      int                 `orm:"column(max_skip_num)"`
	CanGetSkipOrder int8                `orm:"column(can_get_skip_order)"`
	PurifyMissRate  float64             `orm:"column(purify_miss_rate);null;digits(7);decimals(4)"`
	PurifyErrorRate float64             `orm:"column(purify_error_rate);null;digits(7);decimals(4)"`
	DepartmentId    *MainDepartment     `orm:"column(department_id);rel(fk)"`
	PositionId      *MainPositionconfig `orm:"column(position_id);rel(fk)"`
	UserId          *AuthUser           `orm:"column(user_id);rel(fk)"`
}

func (t *MainUserprofile) TableName() string {
	return "main_userprofile"
}

func init() {
	orm.RegisterModel(new(MainUserprofile))
}

// AddMainUserprofile insert a new MainUserprofile into database and returns
// last inserted Id on success.
func AddMainUserprofile(m *MainUserprofile) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetMainUserprofileById retrieves MainUserprofile by Id. Returns error if
// Id doesn't exist
func GetMainUserprofileById(id int) (v *MainUserprofile, err error) {
	o := orm.NewOrm()
	v = &MainUserprofile{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllMainUserprofile retrieves all MainUserprofile matches certain condition. Returns empty list if
// no records exist
func GetAllMainUserprofile(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(MainUserprofile))
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

	var l []MainUserprofile
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

// UpdateMainUserprofile updates MainUserprofile by Id and returns error if
// the record to be updated doesn't exist
func UpdateMainUserprofileById(m *MainUserprofile) (err error) {
	o := orm.NewOrm()
	v := MainUserprofile{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteMainUserprofile deletes MainUserprofile by Id and returns error if
// the record to be deleted doesn't exist
func DeleteMainUserprofile(id int) (err error) {
	o := orm.NewOrm()
	v := MainUserprofile{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&MainUserprofile{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
