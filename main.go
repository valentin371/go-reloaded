package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// 1. Vérifier les arguments
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <input.txt> <output.txt>")
		return
	}
	inputFile := os.Args[1]
	outputFile := os.Args[2]

	// 2. Lire le fichier d'entrée
	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Erreur de lecture :", err)
		return
	}

	// 3. Transformer le texte
	text := string(data)
	words := strings.Fields(text)
	words = ApplyCase(words)
	result := strings.Join(words, " ")

	// 4. Écrire le résultat
	err = os.WriteFile(outputFile, []byte(result), 0644)
	if err != nil {
		fmt.Println("Erreur d'écriture :", err)
		return
	}
}
