package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Booting up webserver...")

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "You've requested on URL: %s\n", r.URL.Path)
	})

	r.Handle("/files/", http.FileServer(http.Dir("static/")))

	log.Println("Booted up webserver")
	http.ListenAndServe(":80", r)
}
