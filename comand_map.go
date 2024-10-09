package main

import (
	"errors"
	"fmt"
)

func commandMapf(cfg *config) error {
	client := cfg.pokeapiClient

	locations, err := client.ListLocations(cfg.nextLocationsPage)

	if err != nil {
		return err
	}

	cfg.nextLocationsPage = locations.Next
	cfg.prevLocationsPage = locations.Prev

	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}

	return nil

}

func commandMapb(cfg *config) error {
	client := cfg.pokeapiClient

	if cfg.prevLocationsPage == nil {
		return errors.New("you are on the first page")
	}

	locations, err := client.ListLocations(cfg.prevLocationsPage)

	if err != nil {
		return err
	}

	cfg.nextLocationsPage = locations.Next
	cfg.prevLocationsPage = locations.Prev

	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
