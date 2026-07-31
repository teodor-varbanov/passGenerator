package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type passEncrypt struct {
	Text string `json:"text"`
}

type storedResponse struct {
	PasswordListID string `json:"PasswordListID"`
	Title          string `json:"Title"`
	UserName       string `json:"UserName"`
	Password       string `json:"Password"`
}

func encRequest(customClient *http.Client, secret string, apiKey string) (*http.Response, error) {
	pass := passEncrypt{
		Text: secret,
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

func passStateRequest(customClient *http.Client, secret string, listID string, title string, username string, stateUrl string) (*http.Response, error) {
	stateInfo := storedResponse{
		PasswordListID: listID,
		Title:          title,
		UserName:       username,
		Password:       secret,
	}

	stateUrl = "http://localhost:6969/api/passwords"
	jsonData, err := json.Marshal(stateInfo)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", stateUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+"something")
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
