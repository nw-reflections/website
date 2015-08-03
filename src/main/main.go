package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
	template.CSS()
}

func main() {
	http.HandleFunc("/hello", HelloServer)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
