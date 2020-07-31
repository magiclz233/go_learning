package main

import (
	"github.com/labstack/echo/v4"
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

func getUser(c echo.Context) error {
	// User ID 来自于url "users/:id"
	id, _ := strconv.Atoi(c.Param("id"))
	for _, user := range users {
		if id == user.Id {
			return c.JSON(http.StatusOK, user)
		}
	}
	return c.String(404, "没有找到id:"+strconv.Itoa(id)+"对应的user")
}

func saveUser(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	users = append(users, *u)
	return c.JSON(http.StatusOK, users)
}

func updateUser(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	for i, user := range users {
		if u.Id == user.Id {
			users[i] = *u
			return c.JSON(http.StatusOK, users)
		}
	}
	return c.String(404, "没有找到该id: "+strconv.Itoa(u.Id)+" 对应的用户")
}

func deleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, user := range users {
		if id == user.Id {
			users = append(users[:i], users[i+1:]...)
		}
	}
	return c.JSON(http.StatusOK, users)
}
