package main

import (
	"fmt"
	"net/http"
)

func home_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Site works with Go on server")
}

func contacts_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Contacts our project")
}

func handleRequest() {
	http.HandleFunc("/", home_page) // Происходит отслеживание домашней страницы
	http.HandleFunc("pages/contact.html", contacts_page)
	http.ListenAndServe(":8181", nil)
}

func main() {
	handleRequest()
}
