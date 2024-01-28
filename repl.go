package main

import (
	"bufio"
	"fmt"
	"github.com/langer-net/pokedexcli/internal/pokeapi"
	"os"
)

type config struct {
	pokeapiClient        pokeapi.Client
	nextLocationsURL     *string
	previousLocationsURL *string
	caughtPokemon        map[string]pokeapi.PokemonResponse
}

func startRepl(cfg *config) {
	cliCommands := getCliCommands()
	cliReader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		cliReader.Scan()

		words := cleanInput(cliReader.Text())
		if len(words) == 0 {
			fmt.Println("Error: no input was given")
			continue
		}
		commandName := words[0]
		var args []string
		if len(words) > 1 {
			args = words[1:]
		}

		command, exists := cliCommands[commandName]
		if !exists {
			fmt.Println("Error: ", commandName, " is not a command")
			fmt.Println("Use help to get a list of all commands.")
			continue
		}
		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
	}
}
