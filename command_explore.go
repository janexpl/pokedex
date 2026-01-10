package main

import "fmt"

func commandExplore(cfg *config, ext ...string) error {
	if len(ext) == 0 {
		return fmt.Errorf("please provide a location name or ID")
	}
	locationPokemons, err := cfg.pokeClient.GetLocation(ext[0])
	if err != nil {
		return err
	}
	for _, pokemon := range locationPokemons.PokemonEncounters {
		fmt.Println(" -", pokemon.Pokemon.Name)
	}
	return nil
}
