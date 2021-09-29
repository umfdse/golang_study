package routers

import (
	"beego-project/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/prev", &controllers.UserController{})
	beego.Router("/pv/:id", &controllers.PagegetController{})
	beego.Router("/page/:url", &controllers.PageController{})
	beego.Router("/postData", &controllers.PagepostController{})
	beego.Router("/postjson", &controllers.PostjsonController{})
	beego.Router("/flash_data", &controllers.Flash_dataController{})
	beego.Router("/getMethods_fn", &controllers.MocsallController{}, "Get,Post:PostFunc")
	beego.Router("/mysql_crud_sql", &controllers.OrmdemoController{})
	beego.Router("/primal_mysql_crud_sql", &controllers.OrmPrimController{})
}
