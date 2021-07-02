package main

import (
	"io"
	"net/http"
	"os"

	about "github.com/panxomon/letargico/about/domain"
	infrastructure "github.com/panxomon/letargico/about/infrastructure"
	server "github.com/panxomon/letargico/infrastructure"
)

const defaultPort = "3001"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	repo, _ := infrastructure.NewDirRepository()
	service := about.NewAboutService(repo)
	aboutHandler := infrastructure.NewAboutHandler(service)

	http.HandleFunc("/home", HelloLucho)
	http.HandleFunc("/about", aboutHandler.Get)

	server.Start(port)

}

func HelloLucho(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "lucho sapbe ingles, server go funcionando")
}
