package main

import (
	"fmt"
	"gnr"
	"net/http"
)

func main() {
	// 需要在go mod中导入gnr包
	r := gnr.New()

	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		_, _ = fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	})

	r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			_, _ = fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})
	_ = r.Run(":8080")
}
