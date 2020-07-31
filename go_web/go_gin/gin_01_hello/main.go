package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World")
	})

	// : 跟的 会匹配 /user/magic 不会匹配 /user 或者 /user/
	r.GET("/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	r.GET("/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		fmt.Println(name, action)
		action = strings.Trim(action, "/")
		msg := name + "is" + action
		c.String(http.StatusOK, msg)
	})
	if err := r.Run(); err != nil {
		fmt.Println(err.Error())
	}
}
