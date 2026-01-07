package main

import "fmt"

func commandExplore(cfg *config, ext ...string) error {
	locationPokemons, err := cfg.pokeClient.ListPokemons(ext[0])
	if err != nil {
		return err
	}
	for _, pokemon := range locationPokemons.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)

	}
	return nil
}
