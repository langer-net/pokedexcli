package poke_api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationAreas struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type LocationAreaRequest struct {
	Count    int             `json:"count"`
	Next     string          `json:"next"`
	Previous string          `json:"previous"`
	Results  []LocationAreas `json:"results"`
}

func makeLocationAreaRequest() (LocationAreaRequest, error) {
	locationAreaRequest := LocationAreaRequest{}

	resp, err := http.Get("https://pokeapi.co/api/v2/location-area/")
	if err != nil {
		return locationAreaRequest, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return locationAreaRequest, fmt.Errorf("failed to read response body: %w", err)
	}

	err = resp.Body.Close()
	if err != nil {
		return locationAreaRequest, fmt.Errorf("failed to close response body: %w", err)
	}

	err = json.Unmarshal(body, &locationAreaRequest)
	if err != nil {
		return locationAreaRequest, fmt.Errorf("failed unmarshal response body: %w", err)
	}
	return locationAreaRequest, nil
}

func ProcessLocationAreaRequest() error {
	locationAreaRequest, err := makeLocationAreaRequest()
	if err != nil {
		return err
	}

	fmt.Println(locationAreaRequest.Results[0])
	fmt.Println(locationAreaRequest.Next)
	fmt.Println(locationAreaRequest.Previous)

	return nil
}
