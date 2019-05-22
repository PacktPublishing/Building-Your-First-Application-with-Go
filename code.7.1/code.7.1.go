package main

import (
	"fmt"
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Hello, world!</h1>")
}

func main() {
	http.HandleFunc("/", RootHandler)
	http.ListenAndServe(":8000", nil)
}
