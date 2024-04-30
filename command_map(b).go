package main

import (
	"fmt"
)

func commandMap(cfg *config, _ string) error {
	maps, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = maps.Next
	cfg.previousLocationsURL = maps.Previous

	for _, result := range maps.Results {
		fmt.Printf("%v\n", result.Name)
	}
	return nil
}

func commandMapB(cfg *config, _ string) error {
	maps, err := cfg.pokeapiClient.ListLocations(cfg.previousLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = maps.Next
	cfg.previousLocationsURL = maps.Previous

	for _, result := range maps.Results {
		fmt.Printf("%v\n", result.Name)
	}
	return nil
}
