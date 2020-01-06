package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["timebook_api/controllers:BookController"] = append(beego.GlobalControllerRouter["timebook_api/controllers:BookController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["timebook_api/controllers:BookController"] = append(beego.GlobalControllerRouter["timebook_api/controllers:BookController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["timebook_api/controllers:BookController"] = append(beego.GlobalControllerRouter["timebook_api/controllers:BookController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["timebook_api/controllers:BookController"] = append(beego.GlobalControllerRouter["timebook_api/controllers:BookController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["timebook_api/controllers:BookController"] = append(beego.GlobalControllerRouter["timebook_api/controllers:BookController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["timebook_api/controllers:BookController"] = append(beego.GlobalControllerRouter["timebook_api/controllers:BookController"],
		beego.ControllerComments{
			Method:           "SendToMe",
			Router:           `/sendMe/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["timebook_api/controllers:BookController"] = append(beego.GlobalControllerRouter["timebook_api/controllers:BookController"],
		beego.ControllerComments{
			Method:           "Update",
			Router:           `/update`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["timebook_api/controllers:TimeController"] = append(beego.GlobalControllerRouter["timebook_api/controllers:TimeController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["timebook_api/controllers:TimeController"] = append(beego.GlobalControllerRouter["timebook_api/controllers:TimeController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["timebook_api/controllers:TimeController"] = append(beego.GlobalControllerRouter["timebook_api/controllers:TimeController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["timebook_api/controllers:TimeController"] = append(beego.GlobalControllerRouter["timebook_api/controllers:TimeController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["timebook_api/controllers:TimeController"] = append(beego.GlobalControllerRouter["timebook_api/controllers:TimeController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["timebook_api/controllers:TimeController"] = append(beego.GlobalControllerRouter["timebook_api/controllers:TimeController"],
		beego.ControllerComments{
			Method:           "GetFreeTime",
			Router:           `/free/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["timebook_api/controllers:TimeController"] = append(beego.GlobalControllerRouter["timebook_api/controllers:TimeController"],
		beego.ControllerComments{
			Method:           "InitTable",
			Router:           `/init/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["timebook_api/controllers:UserController"] = append(beego.GlobalControllerRouter["timebook_api/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["timebook_api/controllers:UserController"] = append(beego.GlobalControllerRouter["timebook_api/controllers:UserController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["timebook_api/controllers:UserController"] = append(beego.GlobalControllerRouter["timebook_api/controllers:UserController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["timebook_api/controllers:UserController"] = append(beego.GlobalControllerRouter["timebook_api/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["timebook_api/controllers:UserController"] = append(beego.GlobalControllerRouter["timebook_api/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["timebook_api/controllers:UserController"] = append(beego.GlobalControllerRouter["timebook_api/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Account",
			Router:           `/account`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["timebook_api/controllers:UserController"] = append(beego.GlobalControllerRouter["timebook_api/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Login",
			Router:           `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
