package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	fullUrl := baseURL + "/pokemon/" + pokemonName

	if entry, ok := c.cache.Get(fullUrl); ok {

		pokemon := Pokemon{}

		err := json.Unmarshal(entry, &pokemon)

		if err != nil {
			return Pokemon{}, err
		}

		return pokemon, nil

	}

	resp, err := http.Get(fullUrl)

	if err != nil {
		return Pokemon{}, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		return Pokemon{}, err
	}

	pokemon := Pokemon{}

	err = json.Unmarshal(data, &pokemon)

	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(fullUrl, data)

	return pokemon, nil
}
