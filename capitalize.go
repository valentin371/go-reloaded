package main


func Capitalize(s string) string {
	var chainerune []rune = []rune(s)
	var nouveaumot bool = true
	for indexrune := 0; indexrune < len(chainerune); indexrune++ {
		c := chainerune[indexrune]
		alphanum := (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
		if alphanum {
			if nouveaumot == true {
				if c >= 'a' && c <= 'z' {
					chainerune[indexrune] = c - ('a' - 'A')
				}
				nouveaumot = false
			} else if c >= 'A' && c <= 'Z' {
				chainerune[indexrune] = c + ('a' - 'A')
			}
		} else {
			nouveaumot = true
		}
	}
	return string(chainerune)
}


func Upper(s string) string {
	chainerune := []rune(s)
	for i := 0; i < len(chainerune); i++ {
		if chainerune[i] >= 'a' && chainerune[i] <= 'z' {
			chainerune[i] = chainerune[i] - ('a' - 'A')
		}
	}
	return string(chainerune)
}


func Lower(s string) string {
	chainerune := []rune(s)
	for i := 0; i < len(chainerune); i++ {
		if chainerune[i] >= 'A' && chainerune[i] <= 'Z' {
			chainerune[i] = chainerune[i] + ('a' - 'A')
		}
	}
	return string(chainerune)
}
