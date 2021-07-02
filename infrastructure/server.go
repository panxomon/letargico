package server

import (
	"log"
	"net/http"
)

type server struct {
	*http.Server
}

func Start(port string) {

	server := newServer(port)

	log.Println("iniciando servidor en el puerto: " + port)

	log.Fatal(server.ListenAndServe())

}

func newServer(listening string) *server {
	s := &http.Server{
		Addr: ":" + listening,
	}

	return &server{s}
}
