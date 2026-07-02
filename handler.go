package main

import (
	"fmt"
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

}

type Validate struct {
	Wrg_passWord string
	Wrg_email    string
}

func loginHandler(w http.ResponseWriter, r *http.Request) {

	password := r.FormValue("password")

	email := r.FormValue("email")

	fmt.Println(password)

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

	valid := Validate{
		Wrg_passWord: "",
		Wrg_email:    "",
	}

	comfirm_Password := r.FormValue("comfirm_pass")
	comfirm_Email := r.FormValue("comfirm_mmail")
	//check for 

}





func pageHandler(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("templates/page.html"))

	tmpl.Execute(w, nil)
}
