package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var (
	tpl   *template.Template
	users map[string]string
)

func main() {
	tpl, _ = template.ParseGlob("src/index/*.html")
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/loginauth", loginAuthHandler)
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/registerauth", registerAuthHandler)
	http.ListenAndServe("localhost:8080", nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("******loginHandler running**********")
	tpl.ExecuteTemplate(w, "login.html", nil)
}

func loginAuthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("******loginAuthHandler running**********")
	r.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")
	fmt.Println("username:", username, "password: ", password)

	for user := range users {
		if user == username && users[user] == password {
			fmt.Fprint(w, "you have successfully logged in :)")
			return
		}
	}
	fmt.Fprint(w, "incorrent password")
	tpl.ExecuteTemplate(w, "login.html", "check username and password")
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("******registerHandler running**********")
	tpl.ExecuteTemplate(w, "register.html", nil)
}

func registerAuthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("******registerAuthHandle running**********")
	r.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")

	users[username] = password
}
