package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCliCommands() map[string]cliCommand {
	var cliCommands = make(map[string]cliCommand)
	cliCommands["help"] = cliCommand{
		name:        "help",
		description: "Displays this help message.",
		callback:    commandHelp,
	}
	cliCommands["map"] = cliCommand{
		name:        "map",
		description: "Displays the names of the next 20 location areas in the world of Pokemon.",
		callback:    commandMapForwards,
	}
	cliCommands["mapb"] = cliCommand{
		name:        "mapb",
		description: "Displays the names of the previous 20 location areas in the world of Pokemon.",
		callback:    commandMapBackwards,
	}
	cliCommands["explore"] = cliCommand{
		name:        "explore <location_name>",
		description: "Explores the given location.",
		callback:    commandExplore,
	}
	cliCommands["catch"] = cliCommand{
		name:        "catch <pokemon_name>",
		description: "Attempts to catch the given pokemon.",
		callback:    commandCatch,
	}
	cliCommands["exit"] = cliCommand{
		name:        "exit",
		description: "Exits the Pokedex.",
		callback:    commandExit,
	}
	return cliCommands
}

func commandHelp(_ *config, _ ...string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: ")
	fmt.Println()
	cliCommands := getCliCommands()
	for _, command := range cliCommands {
		fmt.Printf("%s: %s \r\n", command.name, command.description)
	}
	fmt.Println()
	return nil
}

func commandMapForwards(cfg *config, _ ...string) error {
	locationAreaRequest, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationAreaRequest.Next
	cfg.previousLocationsURL = locationAreaRequest.Previous

	for _, location := range locationAreaRequest.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapBackwards(cfg *config, _ ...string) error {
	locationAreaResponse, err := cfg.pokeapiClient.ListLocations(cfg.previousLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationAreaResponse.Next
	cfg.previousLocationsURL = locationAreaResponse.Previous

	for _, location := range locationAreaResponse.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	name := args[0]
	location, err := cfg.pokeapiClient.GetLocation(name)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring  %s... \r\n", location.Name)
	fmt.Println("Found Pokemon: ")
	for _, encounter := range location.PokemonEncounters {
		fmt.Printf(" - %s \r\n", encounter.Pokemon.Name)
	}

	return nil
}

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	result := rand.Intn(pokemon.BaseExperience)

	fmt.Printf("Throwing a Pokeball at %s ... \r\n", pokemon.Name)
	if result > 40 {
		fmt.Printf("%s escaped! \r\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s was caught! \r\n", pokemon.Name)

	cfg.caughtPokemon[pokemon.Name] = pokemon

	return nil
}

func commandExit(_ *config, _ ...string) error {
	fmt.Println("Exiting the Pokedex ...")
	os.Exit(0)
	return nil
}
