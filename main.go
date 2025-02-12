package main

import (
	"fmt"
	"net/http"
	"html/template"
)

type Data struct{
	Title string 
	Message string 
} 

var tmpl = template.Must(template.ParseFiles("templates/index.html"))

func homeHandler(w http.ResponseWriter, r *http.Request) {
	data := Data{Title:"Home Page Title", Message:"Welcome Home!"}
	tmpl.ExecuteTemplate(w,"index.html",&data)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the about page!")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the contact page!")
}

func servicesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the services page!")
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/contact", contactHandler)
	http.HandleFunc("/services", servicesHandler)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}