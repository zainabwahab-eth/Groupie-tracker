package main

import (
	"fmt"
	"groupie-tracker/config"
	"groupie-tracker/internal/routers"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Loading api...")
	config.Loader()

	r := routers.Routers()

	fmt.Println("Server starting in 4000")
	err := http.ListenAndServe(":4000", r)
	log.Fatal(err)
}
