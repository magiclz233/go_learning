package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var users = []User{
	{Id: 1,
		Name: "magic",
		Age:  23,
		Sex:  "男",
	},
	{Id: 2,
		Name: "uzi",
		Age:  23,
		Sex:  "男",
	},
	{Id: 3,
		Name: "faker",
		Age:  24,
		Sex:  "女",
	},
}

func getUser(c *gin.Context) {
	// User ID 来自于url "users/:id"
	id, _ := strconv.Atoi(c.Param("id"))
	for _, user := range users {
		if id == user.Id {
			c.JSON(http.StatusOK, user)
		}
	}
	c.String(404, "没有找到该id对应的user")
}

func saveUser(c *gin.Context) {
	user := new(User)
	err := c.Bind(user)
	if err != nil {
		c.String(404, "传入的用户格式错误")
	}
	users = append(users, *user)
	c.JSON(http.StatusOK, users)
}

func updateUser(c *gin.Context) {
	u := new(User)
	if err := c.Bind(u); err != nil {
		c.String(404, "传入的用户格式错误")
	}

	for i, user := range users {
		if u.Id == user.Id {
			users[i] = *u
			c.JSON(http.StatusOK, users)
		}
	}
	c.String(404, "没有找到该id: "+strconv.Itoa(u.Id)+" 对应的用户")
}

func deleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, user := range users {
		if id == user.Id {
			users = append(users[:i], users[i+1:]...)
		}
	}
	c.JSON(http.StatusOK, users)
}
