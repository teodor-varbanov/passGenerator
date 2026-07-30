package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

const ASCII_BEGIN = 33
const ASCII_END = 126

func main() {

	parseFlags()
	exclude := strings.Split(exclude, ",")
	secret := generatePassword(lenght, exclude)

	// encryption request
	res, err := encRequest(customClient(), secret, token)
	if err != nil {
		fmt.Printf("Could not make a request: %s\n", err)
		os.Exit(1)
	}
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}

	//pass state response

	passStateResponse, err := passStateRequest(customClient(), secret, "123", "TITLE!", "DODO", "P@SSW0RD")
	if err != nil {
		fmt.Printf("Could not make a request: %s\n", err)
		os.Exit(1)
	}

	passStateBody, err := io.ReadAll(passStateResponse.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("%s\n", resBody)
	fmt.Printf("%s\n", passStateBody)

}
