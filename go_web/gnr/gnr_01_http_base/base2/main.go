package main

import (
	"fmt"
	"log"
	"net/http"
)

// Engine is the uni handler for all requests
type Engine struct{}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		_, _ = fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			_, _ = fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	default:
		_, _ = fmt.Fprintf(w, "404 NOT FOUND %s\n", req.URL.Path)
	}
}

func main() {
	e := new(Engine)
	log.Fatal(http.ListenAndServe(":8080", e))
}
