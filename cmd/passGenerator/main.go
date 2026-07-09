package main

import (
	"fmt"
)

//to-do change the contants to paramters given at command line or via a parameter file
//to-do implement logic to exclude certain numbers based on acii table
//to-do implement logic to exclude numbers based on provided input
//to-do think about a way to make the distribution of numbers and letters better

const ASCII_BEGIN = 33
const ASCII_END = 126
const LENGHT = 16

func main() {
	var exclusions []string
	exclusions = []string{"&", "$", ":", "*", "}", "^", "@"} //example exclusions - to be parameterized

	fmt.Println(generatePassword(LENGHT, exclusions))
}
