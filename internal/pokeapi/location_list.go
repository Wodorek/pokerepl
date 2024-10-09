package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	fullUrl := baseURL + "/location-area"

	if pageURL != nil {
		fullUrl = *pageURL
	}

	if entry, ok := c.cache.Get(fullUrl); ok {
		locations := RespShallowLocations{}

		err := json.Unmarshal(entry, &locations)

		if err != nil {
			return RespShallowLocations{}, err
		}

		return locations, nil
	}

	resp, err := http.Get(fullUrl)

	if err != nil {
		return RespShallowLocations{}, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		return RespShallowLocations{}, err
	}

	locations := RespShallowLocations{}

	err = json.Unmarshal(data, &locations)

	if err != nil {
		return RespShallowLocations{}, err
	}

	c.cache.Add(fullUrl, data)
	return locations, nil
}
