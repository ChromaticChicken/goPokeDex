package main

import "fmt"

func commandExplore(cfg *config, name string) error {
	locationJson, err := cfg.pokeapiClient.ExploreLocation(name)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, pokemon := range locationJson.PokemonEncounters {
		fmt.Printf("- %v\n", pokemon.Pokemon.Name)
	}
	return nil
}
