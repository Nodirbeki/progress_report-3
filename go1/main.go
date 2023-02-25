package main

import (
	"html/template"
	"net/http"
)

func main() {
	handleRequest()
}

func handleRequest() {

	http.HandleFunc("/", home_page)
	http.ListenAndServe(":8080", nil)
}

func home_page(n http.ResponseWriter, r *http.Request) {
	temp, _ := template.ParseFiles("templates/home_page.html")

	temp.Execute(n, "home")
}
