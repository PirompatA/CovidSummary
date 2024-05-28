package client

import (
	"Lineman_project/entity"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type FetchData interface {
	FetchPatients() ([]entity.Patient, error)
}

type fetchData struct {
	request entity.FromRequest
}

func New() FetchData {
	return &fetchData
}

func (f *fetchData) FetchPatients() ([]entity.Patient, error) {
	// URL of the public API
	url := "https://static.wongnai.com/devinterview/covid-cases.json"

	// Make the GET request
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error making GET request: %v", err)
	}
	defer resp.Body.Close()

	// Check for non-200 status codes
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: received non-200 status code: %d", resp.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	// Parse the JSON response
	var fromRequest entity.FromRequest
	if err := json.Unmarshal(body, &fromRequest); err != nil {
		return nil, fmt.Errorf("error parsing JSON: %v", err)
	}

	return fromRequest.Data, nil
}
