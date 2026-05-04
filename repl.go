package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name string
	description string
	callback func() error
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan() 

		userInput := cleanInput(scanner.Text())
		if len(userInput) == 0 {
			continue
		}

		commandName := userInput[0]
		cmd, exists := getCommands()[commandName]
		if exists {
			err := cmd.callback()
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	parts := strings.Fields(text)
	for i := range parts {
		lowered := strings.ToLower(parts[i])
		parts[i] = lowered
	}

	return parts
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name: "exit",
			description: "Exit the pokedex",
			callback: commandExit,
		},
		"help": {
			name: "help",
			description: "Give you informations about the pokedexcli",
			callback: commandHelp,
		},
	}
}
