package routers

import (
	"groupie-tracker/internal/handlers"
	"net/http"
)

func Routers() *http.ServeMux {
	r := http.NewServeMux()
	r.HandleFunc("GET /{$}", handlers.HomeHandler)
	return r
}
