package main

import (
	"log"
	"net/http"
	"text/template"
)

func main() {
	//http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Bad server error")
	}

}
func about(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.html")
}
func home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.html")
}
func contact(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "contact.html")
}

func renderTemplate(w http.ResponseWriter, tepml string) {
	temp, err := template.ParseFiles("template/" + tepml)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	temp.Execute(w, nil)
}
