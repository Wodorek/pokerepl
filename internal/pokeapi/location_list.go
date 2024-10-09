package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespShallowLocation, error) {
	fullUrl := baseURL + "/location/"

	if pageURL != nil {
		fullUrl = *pageURL
	}

	resp, err := http.Get(fullUrl)

	if err != nil {
		return RespShallowLocation{}, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		return RespShallowLocation{}, err
	}

	locations := RespShallowLocation{}

	err = json.Unmarshal(data, &locations)

	if err != nil {
		return RespShallowLocation{}, err
	}

	return locations, nil
}
