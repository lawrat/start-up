package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type priseContact struct {
	Nom         string `json:"nom"`
	Email       string `json:"email"`
	Telephone   string `json:"telephone"`
	Service     string `json:"service"`
	Commentaire string `json:"commentaire"`
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	t, err := template.ParseFiles("./templates" + tmpl + ".html")
	if err != nil {
		fmt.Fprint(w, "MODELE INTROUVABLE...")
	}
	t.Execute(w, nil)
}
func acceuil(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		renderTemplate(w, "acceuil")
	case "POST":
		nom := r.FormValue("nom")
		email := r.FormValue("email")
		telephone := r.FormValue("telephone")
		serviceAttendu := r.FormValue("serviceAttendu")
		commentaire := r.FormValue("commentaire")

		var data []priseContact

	}
}

func services(w http.ResponseWriter, r *http.Request) {

}

func apropos(w http.ResponseWriter, r *http.Request) {

}

func contact(w http.ResponseWriter, r *http.Request) {

}

func entreprise(w http.ResponseWriter, r *http.Request) {

}
