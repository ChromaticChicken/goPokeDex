package main

import (
	"fmt"
)

func commandPokedex(cfg *config, _ string) error {
	fmt.Printf("Your Pokedex:\n")
	cfg.pokeDex.Lock.RLock()
	defer cfg.pokeDex.Lock.RUnlock()
	for _, pokemon := range cfg.pokeDex.CaughtPokemon {
		fmt.Printf("  - %v\n", pokemon.Name)
	}
	return nil
}
