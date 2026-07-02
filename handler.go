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

var tmpl = template.Must(template.ParseFiles("templates/index.html"))

func homeHandler(w http.ResponseWriter, r *http.Request) {

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

func registrationHandler(w http.ResponseWriter, r *http.Request) {

	password := r.FormValue("password")

	email := r.FormValue("email")

	UserInfo[email] = password
	fmt.Println(UserInfo)

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

type Validate struct {
	Wrg_passWord string
	Wrg_email    string
}

func loginHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {

		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/login.html"))

	comfirm_email := r.FormValue("comfirm_email")
	comfirm_password := r.FormValue("comfirm_password")

	storedPassword, ok := UserInfo[comfirm_email]

	switch {

	case !ok:
		data := Validate{

			Wrg_email:    "incorrect email, enter a correct one!!!",
			Wrg_passWord: "",
		}
		tmpl.Execute(w, data)
		return
	case storedPassword != comfirm_password:
		data := Validate{

			Wrg_email:    "",
			Wrg_passWord: "incorrect password, enter a correct one!!!",
		}
		tmpl.Execute(w, data)
		return
	}

	http.Redirect(w, r, "/page", http.StatusSeeOther)
}

func pageHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {

		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/page.html"))

	tmpl.Execute(w, nil)
}
