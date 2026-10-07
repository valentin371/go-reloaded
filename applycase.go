package main

import (
	"strconv"
	"strings"
)


func ApplyCase(words []string) []string {
	result := []string{}

	for i := 0; i < len(words); i++ {
		w := words[i]
		mode := ""
		n := 1

	
		if strings.HasPrefix(w, "(") {
			name := strings.Trim(w, "(),")
			if name == "up" || name == "low" || name == "cap" {
				if strings.HasSuffix(w, ")") {
				
					mode = name
				} else if strings.HasSuffix(w, ",") && i+1 < len(words) {
					
					num, err := strconv.Atoi(strings.TrimSuffix(words[i+1], ")"))
					if err == nil {
						mode = name
						n = num
						i++ 
					}
				}
			}
		}

		
		if mode == "" {
			result = append(result, w)
			continue
		}

		
		start := len(result) - n
		if start < 0 {
			start = 0 
		}
		for j := start; j < len(result); j++ {
			switch mode {
			case "up":
				result[j] = Upper(result[j])
			case "low":
				result[j] = Lower(result[j])
			case "cap":
				result[j] = Capitalize(result[j])
			}
		}
	}
	return result
}
