package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, name string) error {
	if _, ok := cfg.pokeDex.CaughtPokemon[name]; ok {
		fmt.Printf("You have already caught %v\n", name)
		return nil
	}

	pokemon, err := cfg.pokeapiClient.CatchPokemon(name)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %v...\n", name)
	baseExp := pokemon.BaseExperience

	minDifficulty := 50.0
	maxDifficulty := 340.0

	minPercentChance := 33.3
	maxPercentChance := 5.0

	percentChance := minPercentChance + (maxPercentChance-minPercentChance)*(float64(baseExp)-minDifficulty)/(maxDifficulty-minDifficulty)

	randomNum := rand.Float64() * 100.0
	if randomNum <= percentChance {
		fmt.Printf("%v was caught!\n", name)
		cfg.pokeDex.Lock.Lock()
		cfg.pokeDex.CaughtPokemon[name] = pokemon
		cfg.pokeDex.Lock.Unlock()
	} else {
		fmt.Printf("%v escaped!\n", name)
	}

	return nil
}
