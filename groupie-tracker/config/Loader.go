package config

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Loader() {
	artists := []Artist{}
	relations := RelationsIndex{}

	LoadAPI("https://groupietrackers.herokuapp.com/api/artists", &artists)
	LoadAPI("https://groupietrackers.herokuapp.com/api/artists", &relations)
}


func LoadAPI (url string, model any){

	resp, err := http.Get(url)

	if err != nil {
		log.Fatal("Error fetching data", err)
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(model)

	if err != nil {
		log.Fatal("Error decoding response", err)
	}

	fmt.Println(artists)

}