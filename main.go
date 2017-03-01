package main

import (
	"io"
	"net/http"
	"log"
	"fmt"
)

const Version = "0.0.1"

func indexHandler(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Hello world!\n")
	io.WriteString(res, fmt.Sprintf("Web application version: %s\n", Version))
}

func main() {
	http.HandleFunc("/", indexHandler)
	fmt.Println("Listening on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
