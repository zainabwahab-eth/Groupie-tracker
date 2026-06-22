package main

import (
	"fmt"
	"groupie-tracker/config"
)

func main() {
	fmt.Println("Loading api...")
	config.LoadAPI()
}
