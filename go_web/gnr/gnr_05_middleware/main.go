package main

import (
	"gnr"
	"log"
	"net/http"
	"time"
)

func onlyForV2() gnr.HandlerFunc {
	return func(c *gnr.Context) {
		t := time.Now()
		c.Fail(500, "Internal Server Error")
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	r := gnr.New()
	r.Use(gnr.Logger())

	r.GET("/", func(c *gnr.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gnr</h1>")
	})

	v2 := r.Group("/v2")
	v2.Use(onlyForV2())
	{
		v2.GET("/hello/:name", func(c *gnr.Context) {
			c.String(http.StatusOK, "hello %s, path: %s\n", c.Param("name"), c.Path)
		})
	}
	log.Fatal(r.Run(":8080"))
}
