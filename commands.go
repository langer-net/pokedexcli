package main

import (
	"fmt"
	"os"
)

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
