package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	var eslai []int

	fmt.Println("eslai =  ", eslai)

	arreglo := [5]int{1, 2, 3, 4, 5}

	eslai = arreglo[0:4]

	fmt.Println("eslai =  ", eslai)

	http.HandleFunc("/home", homeHandler)
	// http.HandleFunc("/contacto", comandos.ContactoHandler)
	// http.HandleFunc("/blog", comandos.EscribirPost)
	// http.HandleFunc("/galeria", comandos.GaleriaHandler)

	// log.Fatal(http.ListenAndServe(os.Getenv("PORT"), nil))
	log.Fatal(http.ListenAndServe(":3001", nil))
}

func homeHandler(writer http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(writer, "URL.Path = %q\n", r.URL.Path)
}
