package main

import "strings"

// ApplyPunctuation colle . , ! ? : ; au mot précédent et met un espace après.
func ApplyPunctuation(text string) string {
	isP := func(r rune) bool { return strings.ContainsRune(".,!?:;", r) }
	runes := []rune(text)
	out := []rune{}

	for i := 0; i < len(runes); {
		r := runes[i]
		if !isP(r) {
			out = append(out, r)
			i++
			continue
		}
		for len(out) > 0 && out[len(out)-1] == ' ' {
			out = out[:len(out)-1]
		}
		for i < len(runes) && isP(runes[i]) {
			out = append(out, runes[i])
			i++
		}
		for i < len(runes) && runes[i] == ' ' {
			i++
		}
		if i < len(runes) {
			out = append(out, ' ')
		}
	}
	return string(out)
}