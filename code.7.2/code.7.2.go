package main

import (
	"github.com/gorilla/mux"
	"fmt"
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Hello, world!</h1>")
}

func NameHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Fprintf(w, "<h1>%s, %s!</h1>", params["hello"], params["name"])
}

func main() {
	gmux := mux.NewRouter()
	gmux.HandleFunc("/", RootHandler)
	gmux.HandleFunc("/{hello}/{name}", NameHandler)
	http.ListenAndServe(":8000", gmux)
}
