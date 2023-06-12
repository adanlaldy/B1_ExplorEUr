package server

import (
	"html/template"
	"net/http"
)

func home(w http.ResponseWriter, _ *http.Request) {
	page, _ := template.ParseFiles("./front/template/home.html")
	err := page.ExecuteTemplate(w, "home.html", nil)
	if err != nil {
		return
	}
}

func category(w http.ResponseWriter, _ *http.Request) {
	page, _ := template.ParseFiles("./front/template/category.html")

	dataTest := []string{"Place", "Tools", "Information", "+"}

	err := page.ExecuteTemplate(w, "category.html", dataTest)
	if err != nil {
		return
	}
}

func Login(w http.ResponseWriter, _ *http.Request) {
	page, _ := template.ParseFiles("./front/template/login.html")
	err := page.ExecuteTemplate(w, "login.html", nil)
	if err != nil {
		return
	}
}

func Registration(w http.ResponseWriter, _ *http.Request) {
	page, _ := template.ParseFiles("./front/template/registration.html")
	err := page.ExecuteTemplate(w, "registration.html", nil)
	if err != nil {
		return
	}
}

func Chat(w http.ResponseWriter, _ *http.Request) {
	page, _ := template.ParseFiles("./front/template/chat.html")
	err := page.ExecuteTemplate(w, "chat.html", nil)
	if err != nil {
		return
	}
}
