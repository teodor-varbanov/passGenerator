package main

import (
	"flag"
)

var lenght int
var token string
var url string
var exclude string
var username string
var computername string
var localAdminUser string
var localAdminPassword string

//to-do: rework and use cfg. 6 params are needed at least and that's a bit much. 

func flags() {
	flag.IntVar(&lenght, "lenght", 16, "Defines the lenght of the randomly generated password.")
	flag.StringVar(&token, "token", "", "A Bamboo generated token for authentication.")
	flag.StringVar(&url, "url", "http://localhost:6969/rest/api/latest/encrypt", "The URL of the API endpoint that will encrypt the randomly generated password.")
	flag.StringVar(&exclude, "exclude", "", "A list of characters or letters to be excluded from randomly generated password.")
	flag.StringVar(&username, "username", "", "The name of the user that will have their password changed.")
	flag.StringVar(&computername, "ComputerName", "", "The name of the remote server")
	flag.StringVar(&localAdminUser, "LocalAdmin", "administrator", "The local admin that has rights to change the user password.")
	flag.StringVar(&localAdminPassword, "LocalAdminPassword", "", "The password of the local admin.")
}

func parseFlags() {
	flags()
	flag.Parse()
}
