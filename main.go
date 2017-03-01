package main

import (
	"io"
	"net/http"
	"log"
	"fmt"
)

var Version = "onbuild"

func indexHandler(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Hello world!\n")
	io.WriteString(res, fmt.Sprintf("Web application version: %s\n", Version))
}

func main() {
	http.HandleFunc("/", indexHandler)
	fmt.Println("Listening on 8080")
	fmt.Printf("Web application version: %s\n", Version)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
