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

type LocationAreaRequest struct {
	Count    int             `json:"count"`
	Next     *string         `json:"next"`
	Previous *string         `json:"previous"`
	Results  []locationAreas `json:"results"`
}

func (c *Client) ListLocations(pageURL *string) (LocationAreaRequest, error) {
	locationAreaRequest := LocationAreaRequest{}

	url := baseURL + "location-area/"
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		err := json.Unmarshal(val, &locationAreaRequest)
		if err != nil {
			return locationAreaRequest, err
		}
		return locationAreaRequest, nil
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

	c.cache.Add(url, data)

	return locationAreaRequest, nil
}
