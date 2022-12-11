package server

import (
	"log"
	"net/http"
)

func Serve(path string, errorChannel chan error) *http.Server {
	fs := http.FileServer(http.Dir(path))
	http.Handle("/", fs)

	log.Print("Serving site at http://localhost:3001")

	server := &http.Server{Addr: ":3001"}

	go func() {
		errorChannel <- server.ListenAndServe()
	}()

	return server
}
