package main

import (
	"fmt"
	"github.com/langer-net/pokedexcli/internal/poke_api"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
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
		callback:    commandMap,
	}
	return cliCommands
}

func commandHelp() error {
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

func commandExit() error {
	fmt.Println("Exiting the Pokedex ...")
	os.Exit(0)
	return nil
}

func commandMap() error {
	err := poke_api.ProcessLocationAreaRequest()
	if err != nil {
		return err
	}
	return nil
}
