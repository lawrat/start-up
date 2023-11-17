package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

const addr = "localhost"
const port = ":8000"
const FileForContact = "contact.json"
const FileForMeet = "meet.json"

type modeleStruct struct {
	Nom         string `json:"nom"`
	Email       string `json:"email"`
	Telephone   string `json:"telephone"`
	Service     string `json:"service"`
	Commentaire string `json:"commentaire"`
}

type modeleRendezVous struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Mail      string `json:"mail"`
	Phone     string `json:"phone"`
	Canal     string `json:"canal"`
	Lieu      string `json:"lieu"`
	Date      string `json:"date"`
	Heure     string `json:"heure"`
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	t, err := template.ParseFiles("./templates/" + tmpl + ".html")
	if err != nil {
		fmt.Fprint(w, "MODELE INTROUVABLE...")
	}
	t.Execute(w, nil)
}

func accueil(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		renderTemplate(w, "accueil")
	case "POST":
		requestType := r.FormValue("request_type")
		switch requestType {
		case "prise-contact":
			client := modeleStruct{
				Nom:         r.FormValue("nom"),
				Email:       r.FormValue("email"),
				Telephone:   r.FormValue("telephone"),
				Service:     r.FormValue("service-attendu"),
				Commentaire: r.FormValue("commentaire"),
			}

			// Vérifier si le fichier JSON existe
			if _, err := os.Stat(FileForContact); os.IsNotExist(err) {
				// Le fichier n'existe pas, initialiser avec un tableau vide
				err := ioutil.WriteFile(FileForContact, []byte("[]"), 0644)
				if err != nil {
					http.Error(w, "Erreur lors de la création du fichier JSON", http.StatusInternalServerError)
					return
				}
			}

			// Lire le contenu actuel du fichier JSON
			contactActuel, err := ioutil.ReadFile(FileForContact)
			if err != nil {
				http.Error(w, "Erreur lors de la lecture du fichier JSON", http.StatusInternalServerError)
				return
			}

			var clients []modeleStruct

			// Si le fichier JSON existe déjà, décodez-le
			err = json.Unmarshal(contactActuel, &clients)
			if err != nil {
				http.Error(w, "Erreur lors du decodage du fichier JSON", http.StatusInternalServerError)
				return
			}

			// Ajouter le nouvel utilisateur à la liste existante
			clients = append(clients, client)

			// Encodez la liste mise à jour en JSON
			clientJSON, err := json.Marshal(clients)
			if err != nil {
				http.Error(w, "Erreur de conversion", http.StatusInternalServerError)
				return
			}

			// Écrire la liste mise à jour dans le fichier JSON
			err = ioutil.WriteFile(FileForContact, clientJSON, 0644)
			if err != nil {
				http.Error(w, "Erreur de conversion", http.StatusInternalServerError)
				return
			}

			fmt.Println("Données du formulaire ajoutées avec succès à", FileForContact)
			http.Redirect(w, r, "/merci", http.StatusSeeOther)
			return
		case "rendez-vous":
			meet := modeleRendezVous{
				Firstname: r.FormValue("firstname"),
				Lastname:  r.FormValue("lastname"),
				Mail:      r.FormValue("mail"),
				Phone:     r.FormValue("phone"),
				Canal:     r.FormValue("canal"),
				Lieu:      r.FormValue("lieu"),
				Date:      r.FormValue("date"),
				Heure:     r.FormValue("heure"),
			}
			// Vérifier si le fichier JSON existe
			if _, err := os.Stat(FileForMeet); os.IsNotExist(err) {
				// Le fichier n'existe pas, initialiser avec un tableau vide
				err := ioutil.WriteFile(FileForMeet, []byte("[]"), 0644)
				if err != nil {
					http.Error(w, "Erreur lors de la création du fichier JSON", http.StatusInternalServerError)
					return
				}
			}

			// Lire le contenu actuel du fichier JSON
			meetActuel, err := ioutil.ReadFile(FileForMeet)
			if err != nil {
				http.Error(w, "Erreur lors de la lecture du fichier JSON", http.StatusInternalServerError)
				return
			}

			var meets []modeleRendezVous

			// Si le fichier JSON existe déjà, décodez-le
			err = json.Unmarshal(meetActuel, &meets)
			if err != nil {
				http.Error(w, "Erreur lors du decodage du fichier JSON", http.StatusInternalServerError)
				return
			}

			// Ajouter le nouvel utilisateur à la liste existante
			meets = append(meets, meet)

			// Encodez la liste mise à jour en JSON
			meetJSON, err := json.Marshal(meets)
			if err != nil {
				http.Error(w, "Erreur de conversion", http.StatusInternalServerError)
				return
			}

			// Écrire la liste mise à jour dans le fichier JSON
			err = ioutil.WriteFile(FileForMeet, meetJSON, 0644)
			if err != nil {
				http.Error(w, "Erreur de conversion", http.StatusInternalServerError)
				return
			}

			fmt.Println("Données du formulaire ajoutées avec succès à", FileForMeet)
			http.Redirect(w, r, "/merci", http.StatusSeeOther)
			return
		default:
			fmt.Fprint(w, "requette non prise en charge")
		}
	default:
		fmt.Fprint(w, "Méthode non prise en charge")
	}
}
func merci(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "merci")
}
func services(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		renderTemplate(w, "services")
	case "POST":
		requestType := r.FormValue("request_type")
		switch requestType {
		case "rendez-vous":
			meet := modeleRendezVous{
				Firstname: r.FormValue("firstname"),
				Lastname:  r.FormValue("lastname"),
				Mail:      r.FormValue("mail"),
				Phone:     r.FormValue("phone"),
				Canal:     r.FormValue("canal"),
				Lieu:      r.FormValue("lieu"),
				Date:      r.FormValue("date"),
				Heure:     r.FormValue("heure"),
			}
			// Vérifier si le fichier JSON existe
			if _, err := os.Stat(FileForMeet); os.IsNotExist(err) {
				// Le fichier n'existe pas, initialiser avec un tableau vide
				err := ioutil.WriteFile(FileForMeet, []byte("[]"), 0644)
				if err != nil {
					http.Error(w, "Erreur lors de la création du fichier JSON", http.StatusInternalServerError)
					return
				}
			}

			// Lire le contenu actuel du fichier JSON
			meetActuel, err := ioutil.ReadFile(FileForMeet)
			if err != nil {
				http.Error(w, "Erreur lors de la lecture du fichier JSON", http.StatusInternalServerError)
				return
			}

			var meets []modeleRendezVous

			// Si le fichier JSON existe déjà, décodez-le
			err = json.Unmarshal(meetActuel, &meets)
			if err != nil {
				http.Error(w, "Erreur lors du decodage du fichier JSON", http.StatusInternalServerError)
				return
			}

			// Ajouter le nouvel utilisateur à la liste existante
			meets = append(meets, meet)

			// Encodez la liste mise à jour en JSON
			meetJSON, err := json.Marshal(meets)
			if err != nil {
				http.Error(w, "Erreur de conversion", http.StatusInternalServerError)
				return
			}

			// Écrire la liste mise à jour dans le fichier JSON
			err = ioutil.WriteFile(FileForMeet, meetJSON, 0644)
			if err != nil {
				http.Error(w, "Erreur de conversion", http.StatusInternalServerError)
				return
			}

			fmt.Println("Données du formulaire ajoutées avec succès à", FileForMeet)
			http.Redirect(w, r, "/merci", http.StatusSeeOther)
			return
		default:
			fmt.Fprint(w, "methode non prise en charge")
		}

	}
}
func apropos(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "a-propos")
}
func entreprise(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "entreprise")
}
func domicile(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "domicile")
}
func contact(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "contact")
}
func main() {
	http.HandleFunc("/", accueil)
	http.HandleFunc("/merci", merci)
	http.HandleFunc("/services", services)
	http.HandleFunc("/a-propos", apropos)
	http.HandleFunc("/entreprise", entreprise)
	http.HandleFunc("/domicile", domicile)
	http.HandleFunc("/contact", contact)

	fmt.Printf("Serveur écoute sur http://%s%s\n", addr, port)
	http.ListenAndServe(port, nil)
}
