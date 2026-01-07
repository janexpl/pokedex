package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/janexpl/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeClient       pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	var ext []string
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}
		cmd := words[0]
		if len(words) > 1 {
			ext = words[1:]
		}
		command, exists := getCommands()[cmd]
		if exists {
			if err := command.callback(cfg, ext...); err != nil {
				fmt.Println("Error executing command:", err)
			}
			continue
		} else {
			fmt.Println("Unknown command:", cmd)
			continue
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	words := strings.Fields(text)
	return words
}

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
			description: "Get the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore <location>",
			description: "Explore a location to see which Pokemons are there",
			callback:    commandExplore,
		},
	}
}
