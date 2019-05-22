package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"fmt"
	"net/http"
)

// Context
type Greeting struct {
	Hello string
	Name  string
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Hello, world!</h1>")
}

func NameHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	
	t := template.New("hello")
	g := Greeting{
		params["hello"],
		params["name"],
	}
	t.Parse(`<h1>{{.Hello}}, {{.Name}}</h1>`)
	t.Execute(w, g)
}

func main() {
	gmux := mux.NewRouter()
	gmux.HandleFunc("/", RootHandler)
	gmux.HandleFunc("/{hello}/{name}", NameHandler)
	http.ListenAndServe(":8000", gmux)
}
