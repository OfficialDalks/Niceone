package main

import (
	"bytes"
	"errors"
	"flag"
	"net/http"
	"net/url"
	"time"
)

// utils.go holds network utils and function helpers

func get(query string) (*http.Response, error) {
	client := http.Client{
		Timeout: 3 * time.Second,
	}

	req, err := http.NewRequest("GET", query, nil)

	if err != nil {
		return &http.Response{}, err
	}

	return client.Do(req)
}

func post(query string, body []byte) (*http.Response, error) {
	client := http.Client{
		Timeout: 3 * time.Second,
	}

	req, err := http.NewRequest("POST", query, bytes.NewBuffer(body))

	if err != nil {
		return &http.Response{}, err
	}

	req.Header.Set("Content-type", "application/json")

	return client.Do(req)
}

func encodeURL(str string) (string, error) {
	u, err := url.Parse(str)

	if err != nil {
		return "", err
	}

	return u.String(), nil
}

// getCredentials grabs apikeys and auth tokens via flags or environment vars
// prioritizes flags
func getCredentials() (serviceCredentials, error) {
	// TODO: implement environment vars

	credentials := serviceCredentials{}

	flag.StringVar(&credentials.shart.token, "token", "", "token used for bot authentication")
	flag.StringVar(&credentials.radarr.url, "radarr-url", "", "url that points to your radarr app")
	flag.StringVar(&credentials.radarr.apiKey, "radarr-key", "", "api key used for radarr")
	flag.StringVar(&credentials.sonarr.url, "sonarr-url", "", "url that points to your sonarr app")
	flag.StringVar(&credentials.sonarr.url, "sonarr-key", "", "api key used for sonarr")
	flag.BoolVar(&isVerbose, "verbose", false, "output more inforation")

	flag.Parse()

	if credentials.shart.token == "" {
		return credentials, errors.New("a token is required")
	}

	return credentials, nil
}
