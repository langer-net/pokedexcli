package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCliCommands() map[string]cliCommand {
	var cliCommands = make(map[string]cliCommand)
	cliCommands["help"] = cliCommand{
		name:        "help",
		description: "Displays this help message.",
		callback:    commandHelp,
	}
	cliCommands["exit"] = cliCommand{
		name:        "exit",
		description: "Exits the Pokedex.",
		callback:    commandExit,
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
	return cliCommands
}

func commandHelp(cfg *config) error {
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

func commandExit(cfg *config) error {
	fmt.Println("Exiting the Pokedex ...")
	os.Exit(0)
	return nil
}

func commandMapForwards(cfg *config) error {
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

func commandMapBackwards(cfg *config) error {
	locationAreaRequest, err := cfg.pokeapiClient.ListLocations(cfg.previousLocationsURL)
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
