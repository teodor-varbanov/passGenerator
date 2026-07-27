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

type encryptedResponse struct {
	EncryptedText string `json:"encryptedText"`
}

// PasswordListID  = $PasswordListID
// Title           = "North American Core Router 1"
// UserName        = "narouter1"
// Password        = "StenS-Lun#3$2^yc"

type storedResponse struct {
	PasswordListID string `json:"PasswordListID"`
	Title          string `json:"Title"`
	Username       string `json:"Username"`
	Password       string `json:"Password"`
}

func encryptString(w http.ResponseWriter, r *http.Request) {
	var password encryptedPassword

	if err := json.NewDecoder(r.Body).Decode(&password); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	hexPass := sha256.Sum256([]byte(password.Password))
	encPass := hex.EncodeToString(hexPass[:])
	response := encryptedResponse{
		EncryptedText: "BAMFAKE@" + encPass,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func storePassword(w http.ResponseWriter, r *http.Request) {
	var storeInfo storedResponse
	var passwordStore = map[string]storedResponse{}

	if err := json.NewDecoder(r.Body).Decode(&storeInfo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	passwordStore[storeInfo.PasswordListID] = storeInfo

	fmt.Fprintln(w, storeInfo.Password)

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
