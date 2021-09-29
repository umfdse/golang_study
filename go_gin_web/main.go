package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

type LoginForm struct {
	User     string `form:"user" binding:"required" json:"name"`
	Password string `form:"password" binding:"required" json:"password"`
}

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()
	// r.LoadHTMLGlob("view/*") /**/
	// Ping test

	r.Static("/static", "./static")
	r.GET("/favicon.ico", func(c *gin.Context) {
		c.String(http.StatusOK, "")
	})
	r.GET("login", func(c *gin.Context) {
		r.LoadHTMLFiles("./view/p1/Login.tmpl")
		c.HTML(http.StatusOK, "Login", gin.H{})
	})
	r.POST("/login", func(c *gin.Context) {
		// 你可以使用显式绑定声明绑定 multipart form：
		// c.ShouldBindWith(&form, binding.Form)
		// 或者简单地使用 ShouldBind 方法自动绑定：
		var form LoginForm
		// 在这种情况下，将自动选择合适的绑定
		err := c.ShouldBind(&form)
		if err == nil {
			c.JSON(http.StatusOK, form)
		} else {
			fmt.Println("出错啦！！！")
		}
	})
	r.GET("/form_post", func(c *gin.Context) {
		r.LoadHTMLFiles("./view/p1/form01.tmpl")
		c.HTML(http.StatusOK, "form01.tmpl", nil)
	})
	r.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})

	r.GET("/JSONP", func(c *gin.Context) {
		data := map[string]interface{}{
			"foo": "bar",
		}
		c.JSONP(http.StatusOK, data)
	})
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/html", func(c *gin.Context) {
		r.LoadHTMLFiles("./view/demo01.tmpl")
		//r.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
		c.HTML(http.StatusOK, "demo01.tmpl", gin.H{
			"title":   "我是json.title",
			"message": "我是message数据内容",
		})
	})
	r.GET("/html01", func(c *gin.Context) {
		r.LoadHTMLFiles("./view/p1/demo02.tmpl")
		c.HTML(http.StatusOK, "demo02.tmpl", gin.H{
			"msg":     201,
			"title":   "深层目录模板渲染",
			"message": "我是模板文件view/p1/demo02.tmpl",
		})
	})
	r.GET("/prev", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg":      201,
			"content":  "success",
			"respData": [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
		})
	})
	r.GET("/someJSON", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "GO语言",
			"tag":  "<br>",
		}

		// 输出 : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
		c.AsciiJSON(http.StatusOK, data)
	})
	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})
	r.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})
	r.GET("/purejson", func(c *gin.Context) {
		c.PureJSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})
	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	/* example curl for /admin with basicauth header
	   Zm9vOmJhcg== is base64("foo:bar")

		curl -X POST \
	  	http://localhost:8080/admin \
	  	-H 'authorization: Basic Zm9vOmJhcg==' \
	  	-H 'content-type: application/json' \
	  	-d '{"value":"bar"}'
	*/
	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})
	fmt.Println("----------------------------------------------------------------")
	fmt.Printf("%T\n", r)
	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
