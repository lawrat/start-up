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
const FileForEntreprise = "entreprise.json"
const FileForDomicile = "domicile.json"

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

type modeleEntreprise struct {
	Noom           string   `json:"noom"`
	Preenom        string   `json:"preenom"`
	Maail          string   `json:"maail"`
	ServiceAnti    []string `json:"services[]"`
	Type           []string `json:"options[]"`
	Localisation   string   `json:"localisation"`
	TypeEntreprise string   `json:"type-entreprise"`
	NombreSalle    string   `json:"nbre-salle"`
	NombreBureau   string   `json:"nbre-bureau"`
	Jardin         string   `json:"jardin"`
	Cuisine        string   `json:"cuisine"`
	Piscine        string   `json:"piscine"`
	Entrepot       string   `json:"entrepot"`
}

type modeleDomicile struct {
	Noom         string   `json:"noom"`
	Preenom      string   `json:"preenom"`
	Maail        string   `json:"maail"`
	ServiceAnti  []string `json:"services[]"`
	Type         []string `json:"options[]"`
	Localisation string   `json:"localisation"`
	Salle        string   `json:"salle"`
	Chambre      string   `json:"chambre"`
	Jardin       string   `json:"jardin"`
	Cuisine      string   `json:"cuisine"`
	Piscine      string   `json:"piscine"`
	Entrepot     string   `json:"entrepot"`
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
			http.Redirect(w, r, "/mercimeet", http.StatusSeeOther)
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

func mercireservation(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "mercireservation")
}
func mercimeet(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "mercimeet")
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
			http.Redirect(w, r, "/mercimeet", http.StatusSeeOther)
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
	switch r.Method {
	case "GET":
		renderTemplate(w, "entreprise")
	case "POST":
		requestType := r.FormValue("request_type")
		switch requestType {
		case "build":
			reserve := modeleEntreprise{
				Noom:           r.FormValue("noom"),
				Preenom:        r.FormValue("preenom"),
				Maail:          r.FormValue("maail"),
				ServiceAnti:    r.Form["services[]"],
				Type:           r.Form["options[]"],
				Localisation:   r.FormValue("localisation"),
				TypeEntreprise: r.FormValue("type-entreprise"),
				NombreSalle:    r.FormValue("nbre-salle"),
				NombreBureau:   r.FormValue("nbre-bureau"),
				Jardin:         r.FormValue("jardin"),
				Cuisine:        r.FormValue("cuisine"),
				Piscine:        r.FormValue("piscine"),
				Entrepot:       r.FormValue("entrepot"),
			}

			// Vérifier si le fichier JSON existe
			if _, err := os.Stat(FileForEntreprise); os.IsNotExist(err) {
				// Le fichier n'existe pas, initialiser avec un tableau vide
				err := ioutil.WriteFile(FileForEntreprise, []byte("[]"), 0644)
				if err != nil {
					http.Error(w, "Erreur lors de la création du fichier JSON", http.StatusInternalServerError)
					return
				}
			}

			// Lire le contenu actuel du fichier JSON
			entrepriseActuel, err := ioutil.ReadFile(FileForEntreprise)
			if err != nil {
				http.Error(w, "Erreur lors de la lecture du fichier JSON", http.StatusInternalServerError)
				return
			}

			var entreprises []modeleEntreprise

			// Si le fichier JSON existe déjà, décodez-le
			err = json.Unmarshal(entrepriseActuel, &entreprises)
			if err != nil {
				http.Error(w, "Erreur lors du decodage du fichier JSON", http.StatusInternalServerError)
				return
			}

			// Ajouter le nouvel utilisateur à la liste existante
			entreprises = append(entreprises, reserve)

			// Encodez la liste mise à jour en JSON
			entrepriseJSON, err := json.Marshal(entreprises)
			if err != nil {
				http.Error(w, "Erreur de conversion", http.StatusInternalServerError)
				return
			}

			// Écrire la liste mise à jour dans le fichier JSON
			err = ioutil.WriteFile(FileForEntreprise, entrepriseJSON, 0644)
			if err != nil {
				http.Error(w, "Erreur de conversion", http.StatusInternalServerError)
				return
			}

			fmt.Println("Données du formulaire ajoutées avec succès à", FileForEntreprise)
			http.Redirect(w, r, "/mercireservation", http.StatusSeeOther)
			return
		default:
			fmt.Fprint(w, "methode non prise en charge")
		}

	}
}
func domicile(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		renderTemplate(w, "domicile")
	case "POST":
		requestType := r.FormValue("request_type")
		switch requestType {
		case "home":
			domi := modeleDomicile{
				Noom:         r.FormValue("noom"),
				Preenom:      r.FormValue("preenom"),
				Maail:        r.FormValue("maail"),
				ServiceAnti:  r.Form["services[]"],
				Type:         r.Form["options[]"],
				Localisation: r.FormValue("localisation"),
				Salle:        r.FormValue("salle"),
				Chambre:      r.FormValue("chambre"),
				Jardin:       r.FormValue("jardin"),
				Cuisine:      r.FormValue("cuisine"),
				Piscine:      r.FormValue("piscine"),
				Entrepot:     r.FormValue("entrepot"),
			}

			// Vérifier si le fichier JSON existe
			if _, err := os.Stat(FileForEntreprise); os.IsNotExist(err) {
				// Le fichier n'existe pas, initialiser avec un tableau vide
				err := ioutil.WriteFile(FileForEntreprise, []byte("[]"), 0644)
				if err != nil {
					http.Error(w, "Erreur lors de la création du fichier JSON", http.StatusInternalServerError)
					return
				}
			}

			// Lire le contenu actuel du fichier JSON
			domicileActuel, err := ioutil.ReadFile(FileForEntreprise)
			if err != nil {
				http.Error(w, "Erreur lors de la lecture du fichier JSON", http.StatusInternalServerError)
				return
			}

			var domiciles []modeleDomicile

			// Si le fichier JSON existe déjà, décodez-le
			err = json.Unmarshal(domicileActuel, &domiciles)
			if err != nil {
				http.Error(w, "Erreur lors du decodage du fichier JSON", http.StatusInternalServerError)
				return
			}

			// Ajouter le nouvel utilisateur à la liste existante
			domiciles = append(domiciles, domi)

			// Encodez la liste mise à jour en JSON
			domicileJSON, err := json.Marshal(domiciles)
			if err != nil {
				http.Error(w, "Erreur de conversion", http.StatusInternalServerError)
				return
			}

			// Écrire la liste mise à jour dans le fichier JSON
			err = ioutil.WriteFile(FileForDomicile, domicileJSON, 0644)
			if err != nil {
				http.Error(w, "Erreur de conversion", http.StatusInternalServerError)
				return
			}

			fmt.Println("Données du formulaire ajoutées avec succès à", FileForDomicile)
			http.Redirect(w, r, "/mercireservation", http.StatusSeeOther)
			return
		default:
			fmt.Fprint(w, "methode non prise en charge")
		}

	}
}
func contact(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "contact")
}

func devis(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "devis")
}
func main() {
	http.HandleFunc("/", accueil)
	http.HandleFunc("/merci", merci)
	http.HandleFunc("/mercireservation", mercireservation)
	http.HandleFunc("/mercimeet", mercimeet)
	http.HandleFunc("/services", services)
	http.HandleFunc("/a-propos", apropos)
	http.HandleFunc("/entreprise", entreprise)
	http.HandleFunc("/domicile", domicile)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/devis", devis)

	fmt.Printf("Serveur écoute sur http://%s%s\n", addr, port)
	http.ListenAndServe(port, nil)
}
