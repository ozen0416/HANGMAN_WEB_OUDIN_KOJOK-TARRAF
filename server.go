package main

import (
	"hmtl/template"
	"net/http"
	"text/template"
)

type Choices struct {
	Choix int
}

type 

func main() {

	fs := http.FileServer(http.Dir("css"))
	http.Handle("/css/style", http.StripPrefix("/css/", fs))

	tmpl := template.Must(template.ParseFiles("index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := Hangman

		tmpl.Execute(w, data)
	})

	http.HandleFunc("/idCard", func(w http.ResponseWriter, r *http.Request) {


		
		tmpl.Execute(w, data)
	}
	http.ListenAndServe(":80", nil)

}
