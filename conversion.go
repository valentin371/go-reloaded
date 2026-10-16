package main
import "strconv"

// ApplyConvert gère (hex) et (bin) : remplace le mot précédent par sa valeur
func ApplyConvert(words []string) []string {
	result := []string{}

	for _, w := range words {
     if w == "(hex)" || w == "(bin)" {
		base := 16
		if w == "(bin)" {
			base = 2

		}
		// s'il y a un mot avant, on le convertit
		if len(result) > 0 {
			last := len(result) - 1
			num, err := strconv.ParseInt(result[last], base, 64)
			if err == nil {
				result[last] = strconv.FormatInt(num, 10)
			}
		}
	continue 
	 }
	result = append(result, w)
	}
	return result 
}
