package main

import (
	"io"
	"net/http"
	"log"
	"fmt"
)

func indexHandler(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Hello world!")
}

func main() {
	http.HandleFunc("/", indexHandler)
	fmt.Println("Listening on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
