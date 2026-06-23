package main


import(
	"net/http"
	"fmt"
)


func main() {

	http.HandleFunc("/", homeHandler)
	fmt.Println("server is now live on port 8080...")
	
}