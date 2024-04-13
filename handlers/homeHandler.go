package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")

	err := t.Execute(w, nil)

	if err != nil {
		log.Printf("Templates error: %v", err)
	}
}