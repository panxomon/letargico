package main

import (
	"io"
	"letargico/about/domain"
	"letargico/about/infrastructure"
	"net/http"
	"os"

	server "letargico/infrastructure"
)

const defaultPort = "3001"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	repo, _ := infrastructure.NewDirRepository()
	service := domain.NewAboutService(repo)
	aboutHandler := infrastructure.NewAboutHandler(service)

	http.HandleFunc("/home", HelloLucho)
	http.HandleFunc("/about", aboutHandler.Get)

	server.Start(port)
}

func HelloLucho(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "lucho sapbe ingles, server go funcionando")
}
