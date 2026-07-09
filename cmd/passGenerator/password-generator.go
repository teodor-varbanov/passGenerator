package main

import "math/rand/v2"

func generateChar(min int, max int) string {

	char := rand.IntN(max-min) + min
	return string(rune(char))

}

func checkExclusions(char string, exclusions []string) bool {
	if len(exclusions) == 0 {
		return false
	} else {
		var counter = 0
		for _, ex := range exclusions {
			if char == ex {
				counter++
			} else {
				continue
			}
		}
		if counter > 0 {
			return true
		} else {
			return false
		}
	}
}

func generatePassword(LENGHT int, exclusions []string) string {

	var password string

	for i := 1; i <= LENGHT; i++ {
		char := generateChar(ASCII_BEGIN, ASCII_END)

		if checkExclusions(char, exclusions) {
			i--
			continue
		} else {
			password += char
		}
	}
	return password
}
