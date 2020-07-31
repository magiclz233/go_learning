package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.POST("/users", saveUser)
	r.GET("/users/:id", getUser)
	r.PUT("/users", updateUser)
	r.DELETE("/users/:id", deleteUser)
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}

}
