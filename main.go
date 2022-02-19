package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	body := `<h1>Rafa tkm++</h1>`
	fmt.Fprintf(w, body)
}

func main() {
	log.Print("Bitch, you better be joking.")

	http.HandleFunc("/", handler)

	port := "8080"

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}