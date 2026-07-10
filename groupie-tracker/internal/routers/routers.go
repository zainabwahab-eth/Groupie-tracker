package routers

import (
	"groupie-tracker/internal/handlers"
	"net/http"
)

func Routers() {
	r := http.NewServeMux()

	r.HandleFunc("GET /{$}", handlers.HomeHandler)
}
