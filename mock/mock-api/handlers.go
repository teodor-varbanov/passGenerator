package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const TOKEN = "very-secret-token"

type encryptedPassword struct {
	Password string `json:"password"`
}

func encryptString(w http.ResponseWriter, r *http.Request) {
	var password encryptedPassword

	if err := json.NewDecoder(r.Body).Decode(&password); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	encPass := sha256.Sum256([]byte(password.Password))

	fmt.Fprintln(w, "Unencrypted password is: ", password.Password)
	fmt.Fprintln(w, "Encrypted password is: "+"BAMSCRT@"+hex.EncodeToString(encPass[:]))
}

func authenticationMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		auth := r.Header.Get("Authorization")

		apiToken := strings.TrimPrefix(auth, "Bearer ")

		if apiToken == "" {
			http.Error(w, "Empty auth token", http.StatusUnauthorized)
			return
		}

		if apiToken != TOKEN {
			http.Error(w, "Invalid Token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)

	}
}
