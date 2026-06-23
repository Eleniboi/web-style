package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", homeHandler)
	fmt.Println("server is now live on port 8080...")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
		return
	}
}
