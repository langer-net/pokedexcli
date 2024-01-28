package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type locationAreas struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type LocationAreaResponse struct {
	Count    int             `json:"count"`
	Next     *string         `json:"next"`
	Previous *string         `json:"previous"`
	Results  []locationAreas `json:"results"`
}

func (c *Client) ListLocations(pageURL *string) (LocationAreaResponse, error) {
	locationAreaResponse := LocationAreaResponse{}

	url := baseURL + "location-area/"
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		err := json.Unmarshal(val, &locationAreaResponse)
		if err != nil {
			return locationAreaResponse, err
		}
		return locationAreaResponse, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return locationAreaResponse, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return locationAreaResponse, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return locationAreaResponse, err
	}

	err = json.Unmarshal(data, &locationAreaResponse)
	if err != nil {
		return locationAreaResponse, err
	}

	c.cache.Add(url, data)

	return locationAreaResponse, nil
}
