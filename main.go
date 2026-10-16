package main

import (
	"fmt"
	"os"
)

func main() {
// Vérifier les arguments
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <input.txt> <output.txt>")
		return
	}
	inputFile := os.Args[1]
	outputFile := os.Args[2]

// Lire le fichier d'entrée
	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Erreur de lecture :", err)
		return
	}

// Transformer le texte
	result := Process(string(data))

// Écrire le résultat
	err = os.WriteFile(outputFile, []byte(result), 0644)
	if err != nil {
		fmt.Println("Erreur d'écriture :", err)
		return
	}
}
