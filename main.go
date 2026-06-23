package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	// fs := http.FileServer(http.Dir("./static/"))
	// http.Handle("/static/", http.StripPrefix("/static", fs))

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/page", pageHandler)
	fmt.Println("server is now live on http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
		return
	}
}
