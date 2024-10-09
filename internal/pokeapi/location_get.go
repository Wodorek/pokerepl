package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) LocationExplore(locationName string) (Location, error) {
	fullUrl := baseURL + "/location-area/" + locationName

	if val, ok := c.cache.Get(fullUrl); ok {
		location := Location{}

		err := json.Unmarshal(val, &location)

		if err != nil {
			return Location{}, nil
		}

		return location, nil
	}

	resp, err := http.Get(fullUrl)

	if err != nil {
		return Location{}, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		return Location{}, err
	}

	location := Location{}

	err = json.Unmarshal(data, &location)

	if err != nil {
		return Location{}, err
	}

	return location, nil
}
