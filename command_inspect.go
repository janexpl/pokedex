package main

import "fmt"

func commandInspect(cfg *config, ext ...string) error {
	if len(ext) == 0 {
		return fmt.Errorf("please provide a pokemon name")
	}
	pokemon, exists := cfg.caughtPokemon[ext[0]]
	if !exists {
		return fmt.Errorf("you have not caught a pokemon named %s", ext[0])
	}
	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Base Experience:", pokemon.BaseExperience)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Abilities:")
	for _, ability := range pokemon.Abilities {
		fmt.Println(" -", ability.Ability.Name)
	}
	return nil
}
