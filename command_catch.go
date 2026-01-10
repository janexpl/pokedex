package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, ext ...string) error {
	if len(ext) == 0 {
		return fmt.Errorf("please provide a pokemon name")
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", ext[0])
	pokemon, err := cfg.pokeClient.GetPokemon(ext[0])
	if err != nil {
		return err
	}
	experience := pokemon.BaseExperience
	r := rand.Intn(100)
	catchChance := experience / 3
	if r < catchChance {
		fmt.Println(pokemon.Name, "was caught!")
		cfg.caughtPokemon[pokemon.Name] = pokemon
	} else {
		fmt.Println(pokemon.Name, "escaped!")
	}
	return nil
}
