// // package main

// // import (
// // 	"encoding/json"
// // 	"fmt"
// // 	"io/ioutil"

// // 	"github.com/360EntSecGroup-Skylar/excelize/v2"
// // )

// // type Model struct {
// // 	Nom         string `json:"nom"`
// // 	Email       string `json:"email"`
// // 	Telephone   string `json:"telephone"`
// // 	Service     string `json:"service"`
// // 	Commentaire string `json:"commentaire"`
// // }

// // func main() {
// // 	// Charger les données JSON depuis un fichier
// // 	data, err := ioutil.ReadFile("data.json")
// // 	if err != nil {
// // 		fmt.Println("Erreur lors de la lecture du fichier JSON:", err)
// // 		return
// // 	}

// // 	var models []Model
// // 	err = json.Unmarshal(data, &models)
// // 	if err != nil {
// // 		fmt.Println("Erreur lors du décodage du fichier JSON:", err)
// // 		return
// // 	}

// // 	// Créer un nouveau fichier Excel
// // 	f := excelize.NewFile()

// // 	// Ajouter une feuille de calcul
// // 	index := f.NewSheet("Sheet1")

// // 	// Ajouter des en-têtes
// // 	headers := map[string]string{
// // 		"A1": "Nom",
// // 		"B1": "Email",
// // 		"C1": "Téléphone",
// // 		"D1": "Service",
// // 		"E1": "Commentaire",
// // 	}

// // 	for cell, header := range headers {
// // 		f.SetCellValue("Sheet1", cell, header)
// // 	}

// // 	// Ajouter les données
// // 	for i, model := range models {
// // 		row := i + 2
// // 		f.SetCellValue("Sheet1", fmt.Sprintf("A%d", row), model.Nom)
// // 		f.SetCellValue("Sheet1", fmt.Sprintf("B%d", row), model.Email)
// // 		f.SetCellValue("Sheet1", fmt.Sprintf("C%d", row), model.Telephone)
// // 		f.SetCellValue("Sheet1", fmt.Sprintf("D%d", row), model.Service)
// // 		f.SetCellValue("Sheet1", fmt.Sprintf("E%d", row), model.Commentaire)
// // 	}

// // 	// Enregistrer le fichier Excel
// // 	if err := f.SaveAs("output.xlsx"); err != nil {
// // 		fmt.Println("Erreur lors de l'enregistrement du fichier Excel:", err)
// // 		return
// // 	}

// // 	fmt.Println("Fichier Excel généré avec succès.")
// // }

// package main

// import (
// 	"github.com/360EntSecGroup-Skylar/excelize/v2"
// )

// func main() {
// 	// Créer un nouveau fichier Excel
// 	f := excelize.NewFile()

// 	// Ajouter une feuille de calcul
// 	index := f.NewSheet("Sheet1")

// 	// Ajouter des données à la feuille de calcul
// 	f.SetCellValue("Sheet1", "A1", "Hello")
// 	f.SetCellValue("Sheet1", "B1", "World!")

// 	// Enregistrer le fichier Excel
// 	if err := f.SaveAs("example.xlsx"); err != nil {
// 		panic(err)
// 	}
// }

package main
