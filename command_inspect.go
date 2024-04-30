package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, name string) error {
	cfg.pokeDex.Lock.RLock()
	defer cfg.pokeDex.Lock.RUnlock()
	if pokemon, ok := cfg.pokeDex.CaughtPokemon[name]; ok {
		fmt.Printf("Name: %v\n", pokemon.Name)
		fmt.Printf("Height: %v\n", pokemon.Height)
		fmt.Printf("Weight: %v\n", pokemon.Weight)
		fmt.Printf("Stats:\n")
		for _, stat := range pokemon.Stats {
			fmt.Printf("  -%v: %v\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Printf("Types:\n")
		for _, pokemonType := range pokemon.Types {
			fmt.Printf("  - %v\n", pokemonType.Type.Name)
		}

		return nil
	} else {
		return errors.New("you have not yet caught " + name)
	}
}
