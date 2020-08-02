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
		c.HTML(http.StatusOK, "<h1>Hello World</h1>")
	})

	r.GET("/user/:name", func(c *gnr.Context) {
		name := c.Param("name")
		user := User{
			Name: "magic",
			Age:  23,
			Sex:  "男",
		}
		if name == user.Name {
			c.JSON(200, user)
		}
		c.String(500, "no user!")
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

	r.GET("/status/*filepath", func(c *gnr.Context) {
		c.JSON(200, gnr.H{"filepath": c.Param("filepath")})
	})

	log.Fatal(r.Run(":8080"))
}
