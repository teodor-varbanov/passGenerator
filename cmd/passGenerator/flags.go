package main

import (
	"flag"
)

var lenght int
var token string
var url string
var exclude string

func flags() {
	flag.IntVar(&lenght, "lenght", 16, "Defines the lenght of the randomly generated password.")
	flag.StringVar(&token, "token", "", "A Bamboo generated token for authentication.")
	flag.StringVar(&url, "url", "http://localhost:6969/rest/api/latest/encrypt", "The URL of the API endpoint that will encrypt the randomly generated password.")
	flag.StringVar(&exclude, "exclude", "", "A list of characters or letters to be excluded from randomly generated password.")
}

func parseFlags() {
	flags()
	flag.Parse()
}
