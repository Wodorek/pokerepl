package main

import "fmt"

func CommandPokedex(cfg *config, args ...string) error {

	caughtPokemon := cfg.caughtPokemon

	if len(caughtPokemon) == 0 {
		fmt.Println("Your pokedex is empty, go catch some pokemon!")
		return nil
	}

	fmt.Println("Your caught pokemon are:")
	for _, pokemon := range caughtPokemon {
		fmt.Printf("  -%s\n", pokemon.Name)
	}
	return nil
}
