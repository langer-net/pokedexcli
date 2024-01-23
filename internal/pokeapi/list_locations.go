package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type LocationAreas struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type LocationAreaRequest struct {
	Count    int             `json:"count"`
	Next     *string         `json:"next"`
	Previous *string         `json:"previous"`
	Results  []LocationAreas `json:"results"`
}

func (c *Client) ListLocations(pageURL *string) (LocationAreaRequest, error) {
	locationAreaRequest := LocationAreaRequest{}

	url := baseURL + "location-area/"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return locationAreaRequest, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return locationAreaRequest, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return locationAreaRequest, err
	}

	err = json.Unmarshal(data, &locationAreaRequest)
	if err != nil {
		return locationAreaRequest, err
	}

	return locationAreaRequest, nil
}
