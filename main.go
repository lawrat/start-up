package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const addr = "localhost"
const port = ":8000"
const FileForContact = "contact.json"
const FileForMeet = "meet.json"
const FileForEntreprise = "entreprise.json"
const FileForDomicile = "domicile.json"

const commentFile = "commentaire.json"
const devisFile = "devis.json"

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
type modelecommentaire struct {
	Username    string `json:"user"`
	Commentaire string `json:"commenter"`
}
type modeleDevis struct {
	Nom     string `json:"nom"`
	Prenom  string `json:"prenom"`
	Email   string `json:"email"`
	Numero  string `json:"numero"`
	Message string `json:"message"`
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	t, err := template.ParseFiles("./templates/" + tmpl + ".html")
	if err != nil {
		fmt.Fprint(w, "MODELE INTROUVABLE...")
		return
	}

	// Vérifier si le fichier JSON existe
	if _, err := os.Stat(commentFile); os.IsNotExist(err) {
		// Si le fichier n'existe pas,creer le fichier vide
		err := ioutil.WriteFile(commentFile, []byte("[]"), 0644)
		if err != nil {
			fmt.Println("Erreur lors de la création du fichier JSON de commentaires :", err)
			http.Error(w, "Erreur lors de la création du fichier JSON de commentaires", http.StatusInternalServerError)
			t.Execute(w, data)
		}
	} else {
		// Charger les commentaires depuis le fichier JSON
		commentActuel, err := ioutil.ReadFile(commentFile)
		if err != nil {
			fmt.Println("Erreur lors de la lecture du fichier JSON de commentaires :", err)
			http.Error(w, "Erreur lors de la lecture du fichier JSON de commentaires", http.StatusInternalServerError)
			return
		}

		var comments []modelecommentaire

		// Si le fichier JSON existe déjà, décodez-le
		err = json.Unmarshal(commentActuel, &comments)
		if err != nil {
			fmt.Println("Erreur lors du décodage du fichier JSON de commentaires :", err)
			http.Error(w, "Erreur lors du décodage du fichier JSON de commentaires", http.StatusInternalServerError)
			return
		}

		// Ajouter les commentaires au contexte de rendu
		data = map[string]interface{}{
			"Comments": comments,
		}

		t.Execute(w, data)

	}
}

func accueil(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		renderTemplate(w, "accueil", "")
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
		case "commentaire":
			commentOnly := modelecommentaire{
				Username:    getUsernameFromRequest(r),
				Commentaire: r.FormValue("commenter"),
			}

			// Vérifier si le fichier JSON existe
			if _, err := os.Stat(commentFile); os.IsNotExist(err) {
				// Le fichier n'existe pas, initialiser avec un tableau vide
				err := ioutil.WriteFile(commentFile, []byte("[]"), 0644)
				if err != nil {
					http.Error(w, "Erreur lors de la création du fichier JSON", http.StatusInternalServerError)
					return
				}
			}

			// Lire le contenu actuel du fichier JSON
			commentActuel, err := ioutil.ReadFile(commentFile)
			if err != nil {
				http.Error(w, "Erreur lors de la lecture du fichier JSON", http.StatusInternalServerError)
				return
			}

			var comments []modelecommentaire

			// Si le fichier JSON existe déjà, décodez-le
			err = json.Unmarshal(commentActuel, &comments)
			if err != nil {
				http.Error(w, "Erreur lors du decodage du fichier JSON", http.StatusInternalServerError)
				return
			}

			// Ajouter le nouvel utilisateur à la liste existante
			comments = append(comments, commentOnly)

			// Encodez la liste mise à jour en JSON
			commentJSON, err := json.Marshal(comments)
			if err != nil {
				http.Error(w, "Erreur de conversion", http.StatusInternalServerError)
				return
			}

			// Écrire la liste mise à jour dans le fichier JSON
			err = ioutil.WriteFile(commentFile, commentJSON, 0644)
			if err != nil {
				http.Error(w, "Erreur de conversion", http.StatusInternalServerError)
				return
			}

			fmt.Println("Commentaire ajouté avec succès à", commentFile)
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		default:
			fmt.Fprint(w, "requette non prise en charge")
		}
	default:
		fmt.Fprint(w, "Méthode non prise en charge")
	}
}
func getUsernameFromRequest(r *http.Request) string {
	// Récupérer l'en-tête "Authorization" de la requête
	authHeader := r.Header.Get("Authorization")

	// Vérifier si l'en-tête est présent
	if authHeader != "" {
		// Exemple : "Basic dXNlcjpwYXNzd29yZA=="
		// Ici, nous supposons un encodage Basic, mais cela dépendra du type d'authentification utilisé
		// Vous devrez peut-être définir votre propre logique d'extraction du nom d'utilisateur
		// en fonction du type d'authentification que vous utilisez
		// Veuillez noter que l'authentification de base (Basic Auth) transmet le nom d'utilisateur et le mot de passe codés en base64
		// Vous devrez décoder cela si c'est ce que vous utilisez
		// Ne pas utiliser l'authentification de base avec des mots de passe non cryptés dans un environnement de production

		// Exemple d'extraction du nom d'utilisateur de l'en-tête Authorization (Base64)
		decoded, err := base64.StdEncoding.DecodeString(authHeader)
		if err == nil {
			credentials := strings.Split(string(decoded), ":")
			if len(credentials) > 0 {
				return credentials[0]
			}
		}
	}

	// Si l'en-tête "Authorization" n'est pas présent ou ne contient pas d'information sur le nom d'utilisateur,
	// vous pouvez explorer d'autres en-têtes HTTP ou d'autres méthodes en fonction de vos besoins.
	// Notez que ces informations ne sont pas fiables et peuvent être absentes en fonction des configurations du client.

	// Retourner "Anonyme" par défaut
	return "Anonyme"
}
func merci(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "merci", "")
}

func mercireservation(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "mercireservation", "")
}
func mercimeet(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "mercimeet", "")
}
func services(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		renderTemplate(w, "services", "")
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
	renderTemplate(w, "a-propos", "")
}
func entreprise(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		renderTemplate(w, "entreprise", "")
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
		renderTemplate(w, "domicile", "")
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
	renderTemplate(w, "contact", "")
}

func devis(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		renderTemplate(w, "devis", "")
	case "POST":

		requestType := r.FormValue("request_type")
		switch requestType {
		case "devis":
			devi := modeleDevis{
				Nom:     r.FormValue("nom"),
				Prenom:  r.FormValue("prenom"),
				Email:   r.FormValue("email"),
				Numero:  r.FormValue("numero"),
				Message: r.FormValue("message"),
			}

			// Vérifier si le fichier JSON existe
			if _, err := os.Stat(devisFile); os.IsNotExist(err) {
				// Le fichier n'existe pas, initialiser avec un tableau vide
				err := ioutil.WriteFile(devisFile, []byte("[]"), 0644)
				if err != nil {
					http.Error(w, "Erreur lors de la création du fichier JSON", http.StatusInternalServerError)
					return
				}
			}

			// Lire le contenu actuel du fichier JSON
			devisActuel, err := ioutil.ReadFile(devisFile)
			if err != nil {
				http.Error(w, "Erreur lors de la lecture du fichier JSON", http.StatusInternalServerError)
				return
			}

			var devis []modeleDevis

			// Si le fichier JSON existe déjà, décodez-le
			err = json.Unmarshal(devisActuel, &devis)
			if err != nil {
				http.Error(w, "Erreur lors du decodage du fichier JSON", http.StatusInternalServerError)
				return
			}

			// Ajouter le nouvel utilisateur à la liste existante
			devis = append(devis, devi)

			// Encodez la liste mise à jour en JSON
			clientJSON, err := json.Marshal(devis)
			if err != nil {
				http.Error(w, "Erreur de conversion", http.StatusInternalServerError)
				return
			}

			// Écrire la liste mise à jour dans le fichier JSON
			err = ioutil.WriteFile(devisFile, clientJSON, 0644)
			if err != nil {
				http.Error(w, "Erreur de conversion", http.StatusInternalServerError)
				return
			}

			fmt.Println("Données du formulaire ajoutées avec succès à", devisFile)
			http.Redirect(w, r, "/merci", http.StatusSeeOther)
			return
		}
	}
}

// func commenter(w http.ResponseWriter, r *http.Request) {
// 	switch r.Method {
// 	case "GET":
// 		renderTemplate(w, "commenter", "")
// 	case "POST":
// 		commentOnly := modelecommentaire{
// 			Username:    r.FormValue("username"),
// 			Commentaire: r.FormValue("commentaire"),
// 		}

// 		// Vérifier si le fichier JSON existe
// 		if _, err := os.Stat(comment); os.IsNotExist(err) {
// 			// Le fichier n'existe pas, initialiser avec un tableau vide
// 			err := ioutil.WriteFile(comment, []byte("[]"), 0644)
// 			if err != nil {
// 				http.Error(w, "Erreur lors de la création du fichier JSON", http.StatusInternalServerError)
// 				return
// 			}
// 		}

// 		// Lire le contenu actuel du fichier JSON
// 		commentActuel, err := ioutil.ReadFile(comment)
// 		if err != nil {
// 			http.Error(w, "Erreur lors de la lecture du fichier JSON", http.StatusInternalServerError)
// 			return
// 		}

// 		var comments []modelecommentaire

// 		// Si le fichier JSON existe déjà, décodez-le
// 		err = json.Unmarshal(commentActuel, &comments)
// 		if err != nil {
// 			http.Error(w, "Erreur lors du décodage du fichier JSON", http.StatusInternalServerError)
// 			return
// 		}

// 		// Ajouter le nouveau commentaire à la liste existante
// 		comments = append(comments, commentOnly)

// 		// Encodez la liste mise à jour en JSON
// 		commentJSON, err := json.Marshal(comments)
// 		if err != nil {
// 			http.Error(w, "Erreur de conversion", http.StatusInternalServerError)
// 			return
// 		}

// 		// Écrire la liste mise à jour dans le fichier JSON
// 		err = ioutil.WriteFile(comment, commentJSON, 0644)
// 		if err != nil {
// 			http.Error(w, "Erreur de conversion", http.StatusInternalServerError)
// 			return
// 		}

// 		fmt.Println("Commentaire ajouté avec succès à", comment)
// 		return
// 	default:
// 		fmt.Fprint(w, "Méthode non prise en charge")
// 	}
// }

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
	// http.HandleFunc("/commenter", commenter)

	fmt.Printf("Serveur écoute sur http://%s%s\n", addr, port)
	http.ListenAndServe(port, nil)
}
