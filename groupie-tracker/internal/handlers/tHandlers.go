package handlers

import (
	"groupie-tracker/config"
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles("artist.html"))

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "artist.html", config.ArtistsData)

	if err != nil {
		http.Error(w, "Error Loading Page", http.StatusInternalServerError)
	}
}
