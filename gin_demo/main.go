package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func sayHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello golang!",
	})
}

func getBook(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Context",
	})
}

//func main() {
//	r := gin.Default()
//	r.LoadHTMLGlob("templates/**/*")
//	r.GET("/posts/index", func(c *gin.Context) {
//		c.HTML(http.StatusOK, "posts/index.html", gin.H{
//			"title": "posts/index",
//		})
//	})
//	r.GET("users/index", func(c *gin.Context) {
//		c.HTML(http.StatusOK, "users/index.html", gin.H{
//			"title": "users/index",
//		})
//	})
//	//r.GET("/hello", sayHello)
//	//
//	//r.GET("/book", getBook)
//	r.Run(":9090")
//}

func main() {
	r := gin.Default()
	//r.LoadHTMLGlob("templates/**/*")
	//r.LoadHTMLFiles("templates/posts/index.html", "templates/users/index.html")
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	r.LoadHTMLFiles("./index.tmpl")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", "<a>hello</a>")
	})
	//r.GET("/posts/index", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "posts/index.html", gin.H{
	//		"title": "posts/index",
	//	})
	//})
	//
	//r.GET("/users/index", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "users/index.html", gin.H{
	//		"title": "users/index",
	//	})
	//})

	r.Run(":8080")
}
