package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/wodorek/pokerepl/internal/pokeapi"
)

type config struct {
	pokeapiClient     pokeapi.Client
	caughtPokemon     map[string]pokeapi.Pokemon
	nextLocationsPage *string
	prevLocationsPage *string
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokerepl > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		cmdName := words[0]
		args := []string{}

		if len(words) > 1 {
			args = words[1:]
		}

		cmd, exists := getCommands()[cmdName]

		if !exists {

			fmt.Println("Unknown command")
			continue
		}

		err := cmd.callback(cfg, args...)

		if err != nil {
			fmt.Println(err)
		}

	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, args ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "closes the cli",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "explore selected location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "try to catch selected pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "inspect a known pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "show a list of known pokemon",
			callback:    CommandPokedex,
		},
	}
}
