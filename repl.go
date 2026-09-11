package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/PoThePanda1/pokedexcli/internal/pokeapi"
)

func startRepl(cliConfig *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		cmd, exists := cliConfig.registry[commandName] // if cmd present exists = true, cmd is correct value. If cmd doesn't exist then exists = false and cmd is 0 value
		if exists {                                    // evalulate
			if err := cmd.callback(cliConfig); err != nil {
				fmt.Println(err)
			}
			continue // after the command is called, after it returned anything to the console, goes to the next user input  (REPL) next loop after printning
		} else { // evaluate: if the cmd doesn't exist, so not in registry
			fmt.Println("Unknown command")
			continue // again the print loop.
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

// This is the registry, the function means we don't have to initialise it and can be used package level scope
func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays Previous 20 locations",
			callback:    commandMapb,
		},
	}

}

type config struct {
	registry      map[string]cliCommand
	pokeapiClient pokeapi.Client
	nextURL       *string
	previousURL   *string
}
