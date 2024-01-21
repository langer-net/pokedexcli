package main

import (
	"bufio"
	"fmt"
	"os"
)

func startRepl() {
	cliCommands := getCliCommands()
	cliReader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("pokedex > ")
		cliReader.Scan()
		input, err := cleanInput(cliReader.Text())
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}

		command, exists := cliCommands[input]
		if !exists {
			fmt.Println("Error: ", input, " is not a command")
			fmt.Println("Use help to get a list of all commands.")
			continue
		}
		err = command.callback()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCliCommands() map[string]cliCommand {
	var cliCommands = make(map[string]cliCommand)
	cliCommands["help"] = cliCommand{
		name:        "help",
		description: "Displays this help message",
		callback:    commandHelp,
	}
	cliCommands["exit"] = cliCommand{
		name:        "exit",
		description: "Exits the Pokedex",
		callback:    commandExit,
	}
	return cliCommands
}
