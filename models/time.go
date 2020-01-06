package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Time struct {
	Id        int       `orm:"column(id);auto"`
	UserId    int       `orm:"column(user_id)"`
	Date      time.Time `orm:"column(date);type(date)"`
	StartTime time.Time `orm:"column(start_time);type(time)"`
	EndTime   time.Time `orm:"column(end_time);type(time)"`
	IsFree    int8      `orm:"column(is_free)"`
	Event     string    `orm:"column(event);size(255);null"`
}

func (t *Time) TableName() string {
	return "time"
}

func init() {
	orm.RegisterModel(new(Time))
}

// AddTime insert a new Time into database and returns
// last inserted Id on success.
func AddTime(m *Time) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetTimeById retrieves Time by Id. Returns error if
// Id doesn't exist
func GetTimeById(id int) (v *Time, err error) {
	o := orm.NewOrm()
	v = &Time{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

func GetTimeByUserId(id string) (m []map[string]interface{}, err error) {
	cond := orm.NewCondition()

	cond1 := cond.And("UserId__iexact", id)
	query := orm.NewOrm().QueryTable(new(Time))
	query = query.SetCond(cond1)
	var times []Time
	_, err = query.Limit(0, 0).All(&times)
	var timesData []map[string]interface{}
	for _, time := range times {
		dt := time.Date.Format("2006-01-02")
		st := time.StartTime.Format("15:04:05")
		et := time.EndTime.Format("15:04:05")
		temp := map[string]interface{}{"date": dt, "starttime": st, "endtime": et, "event": time.Event, "isfree": time.IsFree}
		timesData = append(timesData, temp)
	}
	return timesData, nil

}

func GetFreeTimeByUserId(id int) (m []map[string]interface{}, err error) {
	cond := orm.NewCondition()

	cond1 := cond.And("UserId__iexact", id)
	query := orm.NewOrm().QueryTable(new(Time))
	query = query.SetCond(cond1)
	var times []Time
	_, err = query.Limit(0, 0).All(&times)
	var timesData []map[string]interface{}
	for _, time := range times {
		if time.IsFree == 0 {
			dt := time.Date.Format("2006-01-02")
			st := time.StartTime.Format("15:04:05")
			et := time.EndTime.Format("15:04:05")
			temp := map[string]interface{}{"date": dt, "starttime": st, "endtime": et}
			timesData = append(timesData, temp)
		}
	}
	return timesData, nil

}

// GetAllTime retrieves all Time matches certain condition. Returns empty list if
// no records exist
func GetAllTime(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Time))
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

	var l []Time
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

// UpdateTime updates Time by Id and returns error if
// the record to be updated doesn't exist
func UpdateTimeById(m *Time) (err error) {
	o := orm.NewOrm()
	v := Time{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteTime deletes Time by Id and returns error if
// the record to be deleted doesn't exist
func DeleteTime(id int) (err error) {
	o := orm.NewOrm()
	v := Time{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Time{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
