package main

import (
	"fmt"
	"io"
	"os"
)

//to-do make it so the program, not the mock server returns the correct output, as it will be in production
//to-do clean up and correct variable names
//to-do change the contants to paramters given at command line or via a parameter file
//to-do implement logic to exclude numbers based on provided input

const ASCII_BEGIN = 33
const ASCII_END = 126
const LENGHT = 16
const apiKey = "very-secret-token"

func main() {
	var exclusions []string
	exclusions = []string{"&", "$", ":", "*", "}", "^", "@"} //example exclusions - to be parameterized

	secret := generatePassword(LENGHT, exclusions)

	res, err := encRequest(customClient(), secret, apiKey)
	if err != nil {
		fmt.Printf("Could not make a request: %s\n", err)
		os.Exit(1)
	}
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", resBody)

}
