package main

import (
	"html/template"
	"net/http"
	"image"
	"image/jpeg"
	"wimp/transform"
)

var (
	images = make(map[string]image.Image)
	templates *template.Template
)

func init() {
	templates = template.Must(
		template.ParseFiles("D:/go/src/wimp/static/root.html", "D:/go/src/wimp/static/editor.html"),
	)
}

func HandleRoot(w http.ResponseWriter, r* http.Request) {
	templates.ExecuteTemplate(w, "root.html", nil)
}

func HandleUpload(w http.ResponseWriter, r* http.Request) {
	file, header, _ := r.FormFile("image")
	image, _, _ := image.Decode(file)
	images[header.Filename] = image
	http.Redirect(w, r, "/editor?name=" + header.Filename, 303)
}

func HandleImage(w http.ResponseWriter, r* http.Request) {
	imageName := r.FormValue("name")
	image := images[imageName]
	jpeg.Encode(w, image, &jpeg.Options{Quality: jpeg.DefaultQuality})
}

func HandleEditor(w http.ResponseWriter, r* http.Request) {
	templates.ExecuteTemplate(w, "editor.html", r.FormValue("name"))
}

func HandleTransform(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	transformation := r.FormValue("transformation")
	images[name] = transform.GetTransformer(transformation).Transform(images[name])
	http.Redirect(w, r, "/editor?name=" + name, 303)
}

func main() {
	http.HandleFunc("/", HandleRoot)
	http.HandleFunc("/upload", HandleUpload)
	http.HandleFunc("/image", HandleImage)
	http.HandleFunc("/transform", HandleTransform)
	http.HandleFunc("/editor", HandleEditor)

	http.ListenAndServe(":8000", nil)
}