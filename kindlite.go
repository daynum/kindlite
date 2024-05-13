package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexPage)
	http.HandleFunc("/show", getHighlights)
	fs := http.FileServer(http.Dir("resources"))
	http.Handle("/resources/", http.StripPrefix("/resources/", fs))
	fmt.Println("Server starting at: http://localhost:8000")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
