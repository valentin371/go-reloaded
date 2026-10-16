package main

import "unicode"

// ApplyQuotes colle les apostrophes au texte qu'elles entourent : ' awesome ' -> 'awesome'
func ApplyQuotes(text string) string {
	runes := []rune(text)
	out := []rune{}
	open := true

	for i := 0; i < len(runes); i++ {
		r := runes[i]
		if r != '\'' {
			out = append(out, r)
			continue
		}
		// contraction (don't) : on ne touche à rien
		if i > 0 && i+1 < len(runes) && unicode.IsLetter(runes[i-1]) && unicode.IsLetter(runes[i+1]) {
			out = append(out, r)
			continue
		}
		if open {
			out = append(out, r)
			for i+1 < len(runes) && runes[i+1] == ' ' { // espaces après l'ouvrante
				i++
			}
		} else {
			for len(out) > 0 && out[len(out)-1] == ' ' { // espaces avant la fermante
				out = out[:len(out)-1]
			}
			out = append(out, r)
		}
		open = !open
	}
	return string(out)
}