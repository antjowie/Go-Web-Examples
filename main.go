package main

import (
	"log"

	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoList struct {
	Title string
	Todos []Todo
}

func main() {
	log.Println("Booting up webserver...")
	tmpl := template.Must(template.ParseFiles("layout.html"))

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintf(w, "You've requested on URL: %s\n", r.URL.Path)
		data := TodoList{
			Title: "My todo list",
			Todos: []Todo{
				{"Book 1", true},
				{"Book 2", true},
				{"Book 3", false},
			},
		}
		tmpl.Execute(w, data)
	})

	// blogR := r.PathPrefix("/blog/").Subrouter()

	// blogR.HandleFunc("/page/{page}", func(w http.ResponseWriter, r *http.Request) {
	// 	page := mux.Vars(r)["page"]
	// 	fmt.Fprintf(w, "This is page %s of the blog", page)
	// })

	log.Println("Booted up webserver")
	http.ListenAndServe(":80", r)
}
