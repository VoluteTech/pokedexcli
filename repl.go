package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/VoluteTech/pokedexcli/internal/api"
)

type config struct {
	pokeapiClient api.Client
	nextLocationURL *string
	prevLocationURL *string
}

type cliCommand struct {
	name string
	description string
	callback func(*config, ...string) error
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan() 

		userInput := cleanInput(scanner.Text())
		if len(userInput) == 0 {
			continue
		}

		commandName := userInput[0]
		args := []string{}
		if len(userInput) > 1 {
			args = userInput[1:]
		}

		cmd, exists := getCommands()[commandName]
		if exists {
			err := cmd.callback(cfg, args...)
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
		"map": {
			name: "map",
			description: "List 20 location areas of the pokemon world",
			callback: commandMapf,
		},
		"mapb": {
			name: "mapb",
			description: "List the last 20 location areas of the pokemon world",
			callback: commandMapb,
		},
		"explore": {
			name: "explore <location_name>",
			description: "List all the pokemons of a specific location area",
			callback: explore,
		},
	}
}
