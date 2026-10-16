package main 

// ApplyArticle transforme "a" en "an" si le mot suivant comment par une voyelle ou un h
func ApplyArticle(words []string) []string {
	for i := 0; i < len(words)-1; i++ {
		if words[i] == "a" || words[i] == "A" {
			next := []rune(words[i+1])
			if len(next) == 0 {
				continue

			}
			first := Lower(string(next[0]))
			if first == "a" || first == "e" || first == "i" || first == "o" || first == "u" || first == "h" {
				if words[i] == "a" {
					words[i] = "an"
				} else {
					words[i] = "An"
				}
			}
		}
	}
	return words
}
