package main 


import (

	"net/http"
	"html/template"
	
)


func homeHandler(w http.ResponseWriter, r *http.Request) {

var tmpl = template.Must(template.ParseFiles("templates/index.html"))

	if r.Method != http.MethodGet{
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	err := tmpl.Execute(w, nil)

	if err != nil{
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
}