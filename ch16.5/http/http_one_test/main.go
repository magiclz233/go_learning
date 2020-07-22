package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloServer)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}

func HelloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("调用HelloServer!")
	_, _ = fmt.Fprintf(w, "Hello,"+req.URL.Path[1:])
}
