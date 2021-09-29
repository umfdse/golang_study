package controllers

import (
	"beego-project/models"
	"fmt"
	"html/template"
	"math/rand"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "project"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

type UserController struct {
	beego.Controller
}

func (u *UserController) Get() {
	u.Data["msg"] = 201
	u.Data["Content"] = "success"
	u.TplName = "V1/prev.html"
}

type PagegetController struct {
	beego.Controller
}

/*GetString能获取所有前端传来的参数无论是get请求还是post请求*/
func (p *PagegetController) Get() {
	fmt.Println(`p.Ctx.Input.Param(":id")`)
	a1 := p.Ctx.Input.Param(":id")
	fmt.Println(a1)
	fmt.Println(`p.Ctx.Input.Query("name")`)
	a2 := p.Ctx.Input.Query("name")
	fmt.Println(a2)
	fmt.Println(`p.GetString(":id")`)
	fmt.Println(p.GetString(":id"))
	fmt.Println(`p.Input().Get("name")`)
	fmt.Println(p.Input().Get("name"))
	fmt.Println(`p.Input()`)
	fmt.Println(p.Input())

	p.TplName = "V1/prev.tpl"
}

type PagepostController struct {
	beego.Controller
}

func (p *PagepostController) Post() {
	// fmt.Println(p.GetString("name"))
	// fmt.Println(p.GetString("pwd"))
	// fmt.Println(p.Input())
	// // p.Ctx.Output.Body([]byte("success"))
	body := p.Ctx.Input.RequestBody
	fmt.Println("-------------------------------------")
	fmt.Println(p.Ctx.Request)
	fmt.Printf("类型是:%T 值是：%v  \r\n", body, string(body))
	p.Ctx.Output.Body(body)
	// p.Ctx.WriteString("success")
	// f, h, err := p.GetFile("file")
	// if err != nil {
	// 	fmt.Println("error")
	// }

	// defer f.Close() // 延迟关闭
	// fmt.Println(h.Filename)

	// // 保存文件
	// p.SaveToFile("file", "upload/"+h.Filename) //后一个参数是保存文件目录名, 文件夹必须要事先就已有,
	// // c.Ctx.WriteString("上传成功")
	// /*输出json格式数据*/
	// p.Data["json"] = map[string]interface{}{
	// 	"content": "success",
	// 	"msg":     201,
	// }
	// p.ServeJSON()
	/*输出XML格式数据*/
	// p.Data["xml"] = map[string]interface{}{
	// 	"content": "success",
	// 	"msg":     201,
	// }
	// p.ServeXML()
	/*输出jsonp格式*/
	// p.Data["jsonp"] = map[string]interface{}{
	// 	"msg":     201,
	// 	"content": "success",
	// }
	// p.serveJSONP()
	/*输出yaml*/
	// p.Data["yaml"] = map[string]interface{}{
	// 	"msg":     201,
	// 	"content": "success",
	// }
	// p.serveYAML()
}

type PageController struct {
	beego.Controller
}

func (that *PageController) Get() {
	url := that.GetString(":url")
	fmt.Println("V1/" + url + ".html")
	that.XSRFExpire = 7200 /*设置xsrf过期时间*/
	that.Data["xsrfdata"] = template.HTML(that.XSRFFormHTML())
	that.TplName = "V1/" + url + ".html"
}

type PostjsonController struct {
	beego.Controller
}

func (p *PostjsonController) Post() {
	flash := beego.NewFlash()
	rand.Seed(time.Now().Unix())
	t := rand.Intn(4)
	if t < 1 {
		flash.Error("随机数小于1")
		flash.Store(&p.Controller)
		p.Redirect("/flash_data", 302)
	} else if t >= 1 && t < 2 {
		flash.Notice("随机数介于1与2之间")
		flash.Store(&p.Controller)
		p.Redirect("/flash_data", 302)
	} else {
		flash.Error("随机数大于2")
		flash.Store(&p.Controller)
		p.Redirect("/flash_data", 302)
	}
	// p.Ctx.Output.Body([]byte("success"))
}

type Flash_dataController struct {
	beego.Controller
}

func (f *Flash_dataController) Get() {
	f.Ctx.Output.Body([]byte("success"))
}

type MocsallController struct {
	beego.Controller
}

func (m *MocsallController) PostFunc() {
	fmt.Println(m.Ctx.Request.Method) /*获取请求方法*/
	// m.StopRun()/*可以终止程序(阻止后面的程序执行)*/
	// XSRF
	m.Ctx.Output.Body([]byte(`{msg:201,content:"success"}`))
}

type OrmdemoController struct {
	beego.Controller
}

func (p *OrmdemoController) Get() {
	o := orm.NewOrm()
	o.Using("default")
	// n, err := o.Insert(&models.User{Name: "孙权", Age: 30, Adress: "江东", Job: "社会老奈"}) //插入一条数据
	// n, err := o.InsertMulti(1, []models.User{{Name: "孙坚", Age: 46, Adress: "江东", Job: "恐怖分子"}, {Name: "张琴", Age: 50, Adress: "中国四川", Job: "农民"}})
	// n, err := o.Delete(&models.User{Name: "张琴"}, "Name") /*删除数据*/
	user := &models.User{Name: "孙二狗子"}
	if o.Read(user, "Name") == nil {
		fmt.Println(user)
		user.Name = "孙坚"
		if num, err := o.Update(user, "Name"); err == nil {
			fmt.Println(num)
		}
	}
	// if err != nil {
	// 	p.Ctx.WriteString(fmt.Sprintf("出错啦: %v \n", err))
	// }
	// fmt.Printf("影响行数 %v \n", n)
	p.Ctx.WriteString("Success")
}

type OrmPrimController struct {
	beego.Controller
}

func (that *OrmPrimController) Get() {
	o := orm.NewOrm()
	// var opts []*models.User
	// o.QueryTable("beego_table01").All(&opts) /*查询beego_table01数据表中所有数据*/
	// // o.QueryTable("beego_table01").Filter("id", 33).All(&opts) /*查询beego_table01数据表中id=33的数据*/
	// fmt.Println("查询到的数据-------------------------------------")
	// // fmt.Println(opts)
	// for _, vo := range opts {
	// 	fmt.Println(*vo)
	// }
	// fmt.Println("查询到的数据-------------------------------------")
	var maps []orm.Params
	num, _ := o.Raw("SELECT * FROM beego_table01").Values(&maps)
	if num > 0 {
		for i, rw := range maps {
			fmt.Println(i, rw)
		}
	}
	that.Ctx.Output.Body([]byte("Success"))
}
