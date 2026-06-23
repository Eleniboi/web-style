package main

import (
	"html/template"
	"net/http"
)

type RegDetail struct {
	Email string
}

var UserInfo = map[string]string{}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	var tmpl = template.Must(template.ParseFiles("templates/index.html"))

	if r.Method != http.MethodGet {
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	err := tmpl.Execute(w, nil)

	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}

	// var Reg_email = r.FormValue("email")

	// if Reg_email == "" || !strings.HasSuffix(Reg_email, `@gmail.com`) {

	// 	Data := RegDetail{
	// 		Email: "Address must end with '@gmail.com'",
	// 	}

	// 	tmpl.Execute(w, Data)
	// }

}

func loginHandler(w http.ResponseWriter, r *http.Request) {

	password := r.FormValue("passord")
	email := r.FormValue("email")

	UserInfo[password] = email

	var tmpl = template.Must(template.ParseFiles("templates/login.html"))

	if r.Method != http.MethodPost {
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	err := tmpl.Execute(w, nil)

	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
}
