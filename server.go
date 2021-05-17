package main

import (
	"fmt"
	"net/http"
	"unicode/utf8"
)

func calc_length(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("username")
	name_length := utf8.RuneCountInString(name)
	fmt.Fprintf(w, "Name: %s Length: %d", name, name_length)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/index.html")
	})

	http.HandleFunc("/name_length", calc_length)

	fmt.Println("Server is listening...")
	server.ListenAndServe()
}
