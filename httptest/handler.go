package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Param struct {
	Name string `json:"name"`
}

func HelloHandler(c *gin.Context) {
	var p Param
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "we need a name",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": fmt.Sprintf("hello %s", p.Name),
	})
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/hello", HelloHandler)
	return router
}

func main() {
	r := SetupRouter()

	r.GET("/hello", HelloHandler)

	err := r.Run(":9090")
	if err != nil {
		return
	}
}
