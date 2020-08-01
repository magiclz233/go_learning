package main

import (
	"gnr"
	"log"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  string `json:"sex"`
}

func main() {
	r := gnr.New()
	r.GET("/", func(c *gnr.Context) {
		c.String(http.StatusOK, "Hello World")
	})

	r.GET("/user", func(c *gnr.Context) {
		user := User{
			Name: "magic",
			Age:  23,
			Sex:  "男",
		}
		c.JSON(200, user)
	})
	r.GET("/hello", func(c *gnr.Context) {
		name := c.Query("name")
		c.HTML(200, "<h1>Hello "+name+"</h1>")
	})

	r.POST("/login", func(c *gnr.Context) {
		c.JSON(200, gnr.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password")})
	})

	log.Fatal(r.Run(":8080"))
}
