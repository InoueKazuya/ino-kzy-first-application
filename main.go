package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", handler);
	http.ListenAndServe(":8080", nil)
}

func handler (writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello World\n");
}