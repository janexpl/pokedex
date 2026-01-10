package main

import "fmt"

func commandPokdex(cfg *config, ext ...string) error {
	fmt.Println("Your Pokedex:")
	for name := range cfg.caughtPokemon {
		fmt.Println(" -", name)
	}
	return nil
}
