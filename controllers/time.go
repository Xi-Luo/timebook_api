package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
	"timebook_api/models"

	"github.com/astaxie/beego"
)

// TimeController operations for Time
type TimeController struct {
	beego.Controller
}

// URLMapping ...
func (c *TimeController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("InitTable", c.InitTable)
	c.Mapping("GetFreeTime", c.GetFreeTime)
}

// Post ...
// @Title Post
// @Description create Time
// @Param	body		body 	models.Time	true		"body for Time content"
// @Success 201 {int} models.Time
// @Failure 403 body is empty
// @router / [post]
func (c *TimeController) Post() {
	var data map[string]interface{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &data); err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}
	fmt.Println(data)
	var v models.Time
	var err error

	timeLayout := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Local")
	ds := data["Date"].(string)
	dt, _ := time.ParseInLocation(timeLayout, ds, loc)
	v.Date = dt

	ss := data["StartTime"].(string)
	st, _ := time.ParseInLocation(timeLayout, ss, loc)
	v.StartTime = st

	es := data["EndTime"].(string)
	et, _ := time.ParseInLocation(timeLayout, es, loc)
	v.EndTime = et

	fmt.Println(data["UserId"].(string))
	UserIdStr := data["UserId"].(string)
	v.UserId, err = strconv.Atoi(UserIdStr)
	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}

	v.IsFree = 0
	_, err = models.AddTime(&v)
	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}
	c.Data["json"] = v
	c.ServeJSON()

}

//func (c *TimeController) Post() {
//	var v models.Time
//	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
//		if _, err := models.AddTime(&v); err == nil {
//			c.Ctx.Output.SetStatus(201)
//			c.Data["json"] = v
//		} else {
//			c.Data["json"] = err.Error()
//		}
//	} else {
//		c.Data["json"] = err.Error()
//	}
//	c.ServeJSON()
//}

// InitTable ...
// @Title Get One
// @Description get User by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :id is empty
// @router /init/:id [get]
func (c *TimeController) InitTable() {
	idStr := c.Ctx.Input.Param(":id")
	times, err := models.GetTimeByUserId(idStr)
	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}
	response := make(map[string]interface{})
	response["timesList"] = times
	fmt.Println(response)
	c.Data["json"] = response
	c.ServeJSON()
}

// GetFreeTime ...
// @Title Get One
// @Description get User by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :id is empty
// @router /free/:id [get]
func (c *TimeController) GetFreeTime() {
	idStr := c.Ctx.Input.Param(":id")
	u, err := models.GetUserByUsername(idStr)
	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}
	times, err := models.GetFreeTimeByUserId(u.Id)
	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}
	response := make(map[string]interface{})
	response["timesList"] = times
	fmt.Println(response)
	c.Data["json"] = response
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Time by id
// @Param	body		body 	models.Time	true		"body for Time content"
// @Success 200 {object} models.Time
// @Failure 403 :id is empty
// @router /:id [get]
func (c *TimeController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetTimeById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		//t1 := v.StartTime.Format("15:04:05")
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Time
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Time
// @Failure 403
// @router / [get]
func (c *TimeController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllTime(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Time
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Time	true		"body for Time content"
// @Success 200 {object} models.Time
// @Failure 403 :id is not int
// @router /:id [put]
func (c *TimeController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Time{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateTimeById(&v); err == nil {
			c.Data["json"] = "OK"
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Time
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *TimeController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteTime(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
