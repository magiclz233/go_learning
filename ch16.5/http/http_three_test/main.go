package main

import (
	"io"
	"net/http"
)

const form = `
	<html><body>
        <form action="#" method="post" name="bar">
            <input type="text" name="in" />
            <input type="submit" value="submit"/>
        </form>
    </body></html>
	`

func SimpleServer(w http.ResponseWriter, r *http.Request) {
	_, _ = io.WriteString(w, "<h1>Hello World</h1>")
}

func FormServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	switch r.Method {
	case "GET":
		_, _ = io.WriteString(w, form)
	case "POST":
		_, _ = io.WriteString(w, r.FormValue("in"))
	}

}
func main() {
	http.HandleFunc("/test1", SimpleServer)
	http.HandleFunc("/test2", FormServer)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
