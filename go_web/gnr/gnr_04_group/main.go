package main

import (
	"gnr"
	"log"
	"net/http"
)

func main() {
	r := gnr.New()
	r.GET("/", func(c *gnr.Context) {
		c.HTML(http.StatusOK, "<h1>Hello World</h1>")
	})

	v1 := r.Group("v1")
	{
		v1.GET("/", func(c *gnr.Context) {
			c.HTML(http.StatusOK, "<h1>Hello World</h1>")
		})

		v1.GET("/hello", func(c *gnr.Context) {
			//expect /hello?name=magic
			c.String(http.StatusOK, "hello %s, path: %s\n", c.Query("name"), c.Path)
		})
	}
	v2 := r.Group("v2")
	{
		v2.GET("/hello/:name", func(c *gnr.Context) {
			// expect /hello/magic
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *gnr.Context) {
			c.JSON(http.StatusOK, gnr.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})
	}

	log.Fatal(r.Run(":8080"))
}
