package main

import (
	"sync"
	"time"

	"github.com/ChromaticChicken/goPokeDex/internal/pokeapi"
	"github.com/ChromaticChicken/goPokeDex/internal/pokedex"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	myPokeDex := pokedex.PokeDex{
		CaughtPokemon: map[string]pokedex.Pokemon{},
		Lock:          &sync.RWMutex{},
	}
	cfg := &config{
		pokeapiClient: pokeClient,
		pokeDex:       myPokeDex,
	}
	startRepl(cfg)

}
