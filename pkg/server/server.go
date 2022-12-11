package server

import (
	"log"
	"net/http"
)

func Serve(path string) error {
	fs := http.FileServer(http.Dir(path))
	http.Handle("/", fs)

	log.Print("Serving site at http://localhost:3001")
	return http.ListenAndServe(":3001", nil)
}
