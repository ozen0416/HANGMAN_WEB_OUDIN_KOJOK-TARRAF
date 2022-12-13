package main

import (
	"html/template"
	"net/http"
	"text/template"
)

type User struct {
	Name    string
	Success bool
}

type Choices struct {
	Facile    string
	Normal    string
	Difficile string
}

func main() {

	fs := http.FileServer(http.Dir("css"))
	http.Handle("/css/style", http.StripPrefix("/css/", fs))

	tmpl := template.Must(template.ParseFiles("index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}
		details := User{
			Name:    r.FormValue("username"),
			Success: true,
		}
		tmpl.Execute(w, details)
	})

	http.ListenAndServe(":80", nil)

}
