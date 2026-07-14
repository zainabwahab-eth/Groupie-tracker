package handlers

import (
	"fmt"
	"groupie-tracker/config"
	"html/template"
	"net/http"
	"strconv"
)

var tmpl = template.Must(template.ParseFiles("../templates/artist.html", "../templates/artist-detail.html"))

// var tmpl = template.Must(template.ParseFiles("../templates/artist.html"))

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "artist.html", config.Artists)

	if err != nil {
		http.Error(w, "Error Loading Page", http.StatusInternalServerError)
	}
}

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	idNum, err := strconv.Atoi(id)

	if err != nil {
		http.Error(w, "Invalid Id", http.StatusBadRequest)
	}

	fmt.Println("idNum", idNum)

	artist := *config.ArtistsData[idNum]

	fmt.Println("artist", artist)

	err2 := tmpl.ExecuteTemplate(w, "artist-detail.html", artist)

	if err2 != nil {
		http.Error(w, "Error Loading Page", http.StatusInternalServerError)
	}
}
