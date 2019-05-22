package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

var templates *template.Template

func loadTemplates() {
	templates = template.Must(template.ParseFiles("src/code.7.4/root.html", "src/code.7.4/hello.html"))
}

// Context
type Greeting struct {
	Hello string
	Name  string
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "root.html", nil)
}

func NameHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	g := Greeting{
		params["hello"],
		params["name"],
	}
	templates.ExecuteTemplate(w, "hello.html", g)
}

func init() {
	loadTemplates()
}

func main() {
	gmux := mux.NewRouter()
	gmux.HandleFunc("/", RootHandler)
	gmux.HandleFunc("/{hello}/{name}", NameHandler)
	http.ListenAndServe(":8000", gmux)
}
