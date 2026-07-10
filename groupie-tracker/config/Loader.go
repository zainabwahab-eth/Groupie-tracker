package config

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var ArtistsData = make(map[int]*Artist)

func Loader() {
	var artists []Artist
	var relations RelationsIndex

	LoadAPI("https://groupietrackers.herokuapp.com/api/artists", &artists)
	LoadAPI("https://groupietrackers.herokuapp.com/api/relation", &relations)

	for i := range artists {
		art := artists[i]
		ArtistsData[art.ID] = &art
	}

	for _, rel := range relations.Index {
		if val, ok := ArtistsData[rel.ID]; ok {
			val.Concert = rel.DatesLocations
		}
	}

	fmt.Println("artists", artists)
	fmt.Println("------------------------")
	art := *ArtistsData[1]
	fmt.Println("artistData", art)

}

func LoadAPI[slice *[]Artist | *RelationsIndex](url string, model slice) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Error fetching data", err)
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(model)
	if err != nil {
		log.Fatal("Error decoding response", err)
	}
}
