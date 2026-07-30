package main

import (
	"math/rand/v2"
	"strings"
)

func generateChar(min int, max int) string {

	char := rand.IntN(max-min) + min
	return string(rune(char))

}

func checkExclusions(char string, exclude []string) bool {
	if len(exclude) == 0 {
		return false
	} else {
		var counter = 0
		for _, ex := range exclude {
			if char == strings.TrimSpace(ex) || char == "\"" || char == "'" || char == "\\" {
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

func generatePassword(LENGHT int, exclude []string) string {

	var password string

	for i := 1; i <= LENGHT; i++ {
		char := generateChar(ASCII_BEGIN, ASCII_END)

		if checkExclusions(char, exclude) {
			i--
			continue
		} else {
			password += char
		}
	}
	return password
}
