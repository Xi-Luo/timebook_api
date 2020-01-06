package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Book struct {
	Id          int       `orm:"column(id);auto"`
	UserId      int       `orm:"column(user_id)"`
	UserName    string    `orm:"column(user_name);size(255)"`
	SendToId    int       `orm:"column(send_To_id)"`
	SendToName  string    `orm:"column(send_To_name);size(255)"`
	Date        time.Time `orm:"column(date);type(date)"`
	StartTime   time.Time `orm:"column(start_time);type(time)"`
	EndTime     time.Time `orm:"column(end_time);type(time)"`
	Event       string    `orm:"column(event);size(255)"`
	IsAccepted  int       `orm:"column(is_accepted)"`
	ElseMessage string    `orm:"column(else_message);size(255)"`
}

func (t *Book) TableName() string {
	return "book"
}

func init() {
	orm.RegisterModel(new(Book))
}

// AddBook insert a new Book into database and returns
// last inserted Id on success.
func AddBook(m *Book) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

func GetBookByUserId(id int) (m []map[string]interface{}, err error) {
	cond := orm.NewCondition()
	cond1 := cond.And("UserId__iexact", id)
	query := orm.NewOrm().QueryTable(new(Book))
	query = query.SetCond(cond1)
	var books []Book
	_, err = query.Limit(0, 0).All(&books)
	var booksData []map[string]interface{}
	for _, book := range books {
		dt := book.Date.Format("2006-01-02")
		st := book.StartTime.Format("15:04:05")
		et := book.EndTime.Format("15:04:05")
		var aStr string
		if book.IsAccepted == 0 {
			aStr = "审核中"
		} else {
			if book.IsAccepted == 1 {
				aStr = "已通过"
			} else {
				aStr = "已拒绝"
			}
		}
		temp := map[string]interface{}{"date": dt, "startTime": st, "endTime": et, "username": book.SendToName, "elseMessage": book.ElseMessage, "content": book.Event, "isAccepted": book.IsAccepted, "aStr": aStr}
		booksData = append(booksData, temp)
	}
	return booksData, nil

}

func GetBookByMeId(id int) (m []map[string]interface{}, err error) {
	cond := orm.NewCondition()
	cond1 := cond.And("SendToId__iexact", id)
	query := orm.NewOrm().QueryTable(new(Book))
	query = query.SetCond(cond1)
	var books []Book
	_, err = query.Limit(0, 0).All(&books)
	var booksData []map[string]interface{}
	for _, book := range books {
		if book.IsAccepted == 0 {
			dt := book.Date.Format("2006-01-02")
			st := book.StartTime.Format("15:04:05")
			et := book.EndTime.Format("15:04:05")
			temp := map[string]interface{}{"id": book.Id, "username": book.UserName, "date": dt, "startTime": st, "endTime": et, "content": book.Event}
			booksData = append(booksData, temp)
		}
	}
	return booksData, nil
}

// GetBookById retrieves Book by Id. Returns error if
// Id doesn't exist
func GetBookById(id int) (v *Book, err error) {
	o := orm.NewOrm()
	v = &Book{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllBook retrieves all Book matches certain condition. Returns empty list if
// no records exist
func GetAllBook(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Book))
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

	var l []Book
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

// UpdateBook updates Book by Id and returns error if
// the record to be updated doesn't exist
func UpdateBookById(m *Book) (err error) {
	o := orm.NewOrm()
	v := Book{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteBook deletes Book by Id and returns error if
// the record to be deleted doesn't exist
func DeleteBook(id int) (err error) {
	o := orm.NewOrm()
	v := Book{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Book{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
