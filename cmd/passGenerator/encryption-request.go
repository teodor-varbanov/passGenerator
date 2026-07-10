package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type passEncrypt struct {
	Password string `json:"password"` //here it should probably be a marshalled passwod taken from the pass generator
}

func encRequest(customClient *http.Client, secret string, apiKey string) (*http.Response, error) {
	pass := passEncrypt{
		Password: secret,
	}

	jsonData, err := json.Marshal(pass)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("User-Agent", "Go-HTTP-Client/1.0")

	return customClient.Do(req)
}

func customClient() *http.Client {
	transport := &http.Transport{
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 10,
		IdleConnTimeout:     90 * time.Second,
		DisableKeepAlives:   false,
		DisableCompression:  false,
	}

	return &http.Client{
		Transport: transport,
		Timeout:   30 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			if len(via) >= 10 {
				return errors.New("Stopped after 10 redirects.")
			}
			return nil
		},
	}

}
