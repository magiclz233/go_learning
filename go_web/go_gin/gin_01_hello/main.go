package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func main() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello World")
	})

	// : 跟的 会匹配 /user/magic 不会匹配 /user 或者 /user/
	r.GET("/hello/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	r.GET("/hello/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		fmt.Println(name, action)
		action = strings.Trim(action, "/")
		msg := name + "is" + action
		c.String(http.StatusOK, msg)
	})

	r.GET("/user1", helloUser1)
	r.GET("/user2", helloUser2)

	if err := r.Run(); err != nil {
		fmt.Println(err.Error())
	}
}

// Url形式的  /user?name=123
func helloUser1(c *gin.Context) {
	//默认值自己定义  我定义为了magic
	name := c.DefaultQuery("name", "magic")
	c.String(http.StatusOK, fmt.Sprintf("hello %s", name))
}

// Url形式的  /user?name=123
func helloUser2(c *gin.Context) {
	//默认值为nil
	name := c.Query("name")
	c.String(http.StatusOK, fmt.Sprintf("hello %s", name))
}
