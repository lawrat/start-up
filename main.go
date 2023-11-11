// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"html/template"
// 	"io/ioutil"
// 	"net/http"
// 	"os"
// )

// const addr = "localhost"
// const port = ":8000"
// const jsonFileName = "contact.json"

// type modeleStruct struct {
// 	Nom         string `json:"nom"`
// 	Email       string `json:"email"`
// 	Telephone   string `json:"telephone"`
// 	Service     string `json:"service"`
// 	Commentaire string `json:"commentaire"`
// }

// func renderTemplate(w http.ResponseWriter, tmpl string) {
// 	t, err := template.ParseFiles("./templates/" + tmpl + ".html")
// 	if err != nil {
// 		fmt.Fprint(w, "MODELE INTROUVABLE...")
// 	}
// 	t.Execute(w, nil)
// }

// func accueil(w http.ResponseWriter, r *http.Request) {
// 	switch r.Method {
// 	case "GET":
// 		renderTemplate(w, "accueil")
// 	case "POST":
// 		client := modeleStruct{
// 			Nom:         r.FormValue("nom"),
// 			Email:       r.FormValue("email"),
// 			Telephone:   r.FormValue("telephone"),
// 			Service:     r.FormValue("serviceAttendu"),
// 			Commentaire: r.FormValue("commentaire"),
// 		}

// 		// Vérifier si le fichier JSON existe
// 		if _, err := os.Stat(jsonFileName); os.IsNotExist(err) {
// 			// Le fichier n'existe pas, initialiser avec un tableau vide
// 			err := ioutil.WriteFile(jsonFileName, []byte("[]"), 0644)
// 			if err != nil {
// 				http.Error(w, "Erreur lors de la création du fichier JSON", http.StatusInternalServerError)
// 				return
// 			}
// 		}

// 		// Lire le contenu actuel du fichier JSON
// 		contactActuel, err := ioutil.ReadFile(jsonFileName)
// 		if err != nil {
// 			http.Error(w, "Erreur lors de la lecture du fichier JSON", http.StatusInternalServerError)
// 			return
// 		}

// 		var clients []modeleStruct

// 		// Si le fichier JSON existe déjà, décodez-le
// 		err = json.Unmarshal(contactActuel, &clients)
// 		if err != nil {
// 			http.Error(w, "Erreur lors du decodage du fichier JSON", http.StatusInternalServerError)
// 			return
// 		}

// 		// Ajouter le nouvel utilisateur à la liste existante
// 		clients = append(clients, client)

// 		// Encodez la liste mise à jour en JSON
// 		clientJSON, err := json.Marshal(clients)
// 		if err != nil {
// 			http.Error(w, "Erreur de conversion", http.StatusInternalServerError)
// 			return
// 		}

// 		// Écrire la liste mise à jour dans le fichier JSON
// 		err = ioutil.WriteFile(jsonFileName, clientJSON, 0644)
// 		if err != nil {
// 			http.Error(w, "Erreur de conversion", http.StatusInternalServerError)
// 			return
// 		}

// 		fmt.Println("Données du formulaire ajoutées avec succès à", jsonFileName)
// 		http.Redirect(w, r, "/merci", http.StatusSeeOther)
// 		return
// 	default:
// 		fmt.Fprint(w, "Méthode non prise en charge")
// 	}
// }
// func merci(w http.ResponseWriter, r *http.Request) {
// 	renderTemplate(w, "merci")
// }

// func main() {
// 	http.HandleFunc("/", accueil)
// 	http.HandleFunc("/merci", merci)

// 	fmt.Printf("Serveur écoute sur http://%s%s\n", addr, port)
// 	http.ListenAndServe(port, nil)
// }

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
const port = ":6060"
const fileJson = "data.json"

type modele struct {
	Nom         string `json:"nom"`
	Email       string `json:"email"`
	Telephone   string `json:"telephone"`
	Service     string `json:"service"`
	Commentaire string `json:"commentaire"`
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	t, err := template.ParseFiles("./templates/" + tmpl + ".html")
	if err != nil {
		http.Error(w, "MODELE INTROUVABLE...", http.StatusInternalServerError)
	}
	t.Execute(w, nil)
}

func accueil(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		renderTemplate(w, "accueil")
	case "POST":

		// assigner les données au modele de la structure

		user := modele{
			Nom:         r.FormValue("nom"),
			Email:       r.FormValue("email"),
			Telephone:   r.FormValue("telephone"),
			Service:     r.FormValue("service"),
			Commentaire: r.FormValue("Ccommentaire"),
		}
		// verifier si le fichier json existe
		if _, err := os.Stat(fileJson); os.IsNotExist(err) {
			// initialiser le fichier json avec un tableau vide
			err = os.WriteFile(fileJson, []byte("[]"), 0644)
			if err != nil {
				http.Error(w, "ERREUR LORS DE LA CREATION DU FICHIER JSON", http.StatusInternalServerError)
				return
			}
		}
		// lecture du fichier json
		currentFile, err := ioutil.ReadFile(fileJson)
		if err != nil {
			http.Error(w, "ERREUR LORS DE LA LECTURE DU FICHIER JSON", http.StatusInternalServerError)
		}
		var users []modele
		// decodez le fichier JSON
		err = json.Unmarshal(currentFile, &users)
		if err != nil {
			http.Error(w, "ERREUR LORS DE DECODAGE DU FICHIER JSON", http.StatusInternalServerError)
		}
		// ajouter la nouvelle entre dans le fichier JSON
		users = append(users, user)

		// encodez la liste mis a jour
		newsJson, err := json.Marshal(users)
		if err != nil {
			http.Error(w, "ERREUR LORS DE L'ENCODAGE FDE LA LISTE JSON A JOUR", http.StatusInternalServerError)
		}
		// ecrire la liste mis a jour
		err = ioutil.WriteFile("fileJson", newsJson, 0644)
		if err != nil {
			http.Error(w, "ERREUR LORS DE L'ECRITURE DANS LE FICHIER JSON", http.StatusInternalServerError)
		}
		// imprimer une confirmation d'ecriture
		fmt.Println("DONNEES MIS A JOUR DANS LE FICHIER JSON")
		// rediriger l'utilisateur vers la page de remerciement
		http.Redirect(w, r, "/merci", http.StatusSeeOther)

	default:
		fmt.Fprint(w, "METHODE NON PRIS EN CHARGE")
	}
}
func merci(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "merci")
}

func services(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "services")
}

func apropos(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "a-propos")
}

func contact(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		renderTemplate(w, "contact")
	case "POST":

		// assigner les données au modele de la structure

		user := modele{
			Nom:         r.FormValue("nom"),
			Email:       r.FormValue("email"),
			Telephone:   r.FormValue("telephone"),
			Service:     r.FormValue("service"),
			Commentaire: r.FormValue("Ccommentaire"),
		}
		// verifier si le fichier json existe
		if _, err := os.Stat(fileJson); os.IsNotExist(err) {
			// initialiser le fichier json avec un tableau vide
			err = os.WriteFile(fileJson, []byte("[]"), 0644)
			if err != nil {
				http.Error(w, "ERREUR LORS DE LA CREATION DU FICHIER JSON", http.StatusInternalServerError)
				return
			}
		}
		// lecture du fichier json
		currentFile, err := ioutil.ReadFile(fileJson)
		if err != nil {
			http.Error(w, "ERREUR LORS DE LA LECTURE DU FICHIER JSON", http.StatusInternalServerError)
		}
		var users []modele
		// decodez le fichier JSON
		err = json.Unmarshal(currentFile, &users)
		if err != nil {
			http.Error(w, "ERREUR LORS DE DECODAGE DU FICHIER JSON", http.StatusInternalServerError)
		}
		// ajouter la nouvelle entre dans le fichier JSON
		users = append(users, user)

		// encodez la liste mis a jour
		newsJson, err := json.Marshal(users)
		if err != nil {
			http.Error(w, "ERREUR LORS DE L'ENCODAGE FDE LA LISTE JSON A JOUR", http.StatusInternalServerError)
		}
		// ecrire la liste mis a jour
		err = ioutil.WriteFile("fileJson", newsJson, 0644)
		if err != nil {
			http.Error(w, "ERREUR LORS DE L'ECRITURE DANS LE FICHIER JSON", http.StatusInternalServerError)
		}
		// imprimer une confirmation d'ecriture
		fmt.Println("DONNEES MIS A JOUR DANS LE FICHIER JSON")
		// rediriger l'utilisateur vers la page de remerciement
		http.Redirect(w, r, "/merci", http.StatusSeeOther)

	default:
		fmt.Fprint(w, "METHODE NON PRIS EN CHARGE")
	}
}

func entreprise(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "entreprise")
}
func main() {
	http.HandleFunc("/", accueil)
	http.HandleFunc("/merci", merci)
	http.HandleFunc("/services", services)
	http.HandleFunc("/a-propos", apropos)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/entreprise", entreprise)

	fmt.Printf("serveur en cours d'execution sur http://%s%s\n", addr, port)
	http.ListenAndServe(port, nil)
}
