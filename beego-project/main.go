package main

import (
	_ "beego-project/routers"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	fmt.Println("init函数已执行")
	orm.RegisterDriver("mysql", orm.DRMySQL)                                                                   /*mysql驱动*/
	orm.RegisterDataBase("default", "mysql", "root:Xsb@#123@tcp(127.0.0.1:3306)/beego_sql?charset=utf8", 1000) /*1000为数据库最大链接数量*/

}
func main() {
	// beego.SetStaticPath("/static", "staticDir")  /*设置静态资源文件夹*/
	// beego.SetViewPath("view") /*设置模板指向的文件夹*/
	fmt.Println("---------------------------------------------")
	fmt.Println(beego.AppConfig.String("appname"))
	fmt.Println("---------------------------------------------")
	orm.Debug = true                                        /*开启orm调试模式*/
	beego.SetLogger("file", `{"filename":"logs/test.log"}`) /*开启日志*/
	beego.BConfig.WebConfig.EnableXSRF = true               /*开启XSRF身份验证*/ /*也可在app.conf中设置enablexsrf=true开启*/
	// beego.BConfig.WebConfig.TemplateLeft = "[["             // 设置模板解析变量的起始符号  不设置默认是{{
	// beefo.BConfig.WebConfig.TemplateRight = "]]"            // 设置模板解析变量的终止符号	不设置默认是}}

	beego.Run()
}
