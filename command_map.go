package main

import (
	"errors"
	"fmt"
)

func commandMapf(cfg *config, ext ...string) error {
	locationAreas, err := cfg.pokeClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}
	cfg.nextLocationsURL = locationAreas.Next
	cfg.prevLocationsURL = locationAreas.Previous

	for _, loc := range locationAreas.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(cfg *config, ext ...string) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}
	locationResp, err := cfg.pokeClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}
	cfg.nextLocationsURL = locationResp.Next
	cfg.prevLocationsURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
