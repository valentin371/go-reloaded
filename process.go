package main

import "strings"

// Process applique toutes les règles dans le bon ordre
func Process(text string) string {
	words := strings.Fields(text)
	words = ApplyConvert(words)
	words = ApplyCase(words)
	words = ApplyArticle(words)
	result := strings.Join(words, " ")
	result = ApplyPunctuation(result)
	result = ApplyQuotes(result)
	return result
}
