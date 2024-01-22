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
