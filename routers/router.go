// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"timebook_api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/book",
			beego.NSInclude(
				&controllers.BookController{},
			),
		),

		beego.NSNamespace("/bookreply",
			beego.NSInclude(
				&controllers.BookreplyController{},
			),
		),

		beego.NSNamespace("/time",
			beego.NSInclude(
				&controllers.TimeController{},
			),
		),

		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
