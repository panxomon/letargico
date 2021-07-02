package main

import (
	"io"
	"net/http"
	"os"

	"github.com/panxomon/letargico/about/domain"
	"github.com/panxomon/letargico/about/infraestructure"
	server "github.com/panxomon/letargico/infraestructure"
)

const defaultPort = "3001"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	repo, _ := infraestructure.NewDirRepository()
	service := domain.NewAboutService(repo)
	aboutHandler := infraestructure.NewAboutHandler(service)

	http.HandleFunc("/home", HelloLucho)
	http.HandleFunc("/about", aboutHandler.Get)

	server.Start(port)

}

func HelloLucho(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "lucho sapbe ingles, server go funcionando")
}
