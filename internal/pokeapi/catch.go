package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/ChromaticChicken/goPokeDex/internal/pokedex"
)

func (c *Client) CatchPokemon(name string) (pokedex.Pokemon, error) {
	url := baseURL + "/pokemon/" + name

	if body, ok := c.cache.Get(url); ok {
		exploreJson := pokedex.Pokemon{}
		err := json.Unmarshal(body, &exploreJson)
		if err != nil {
			return pokedex.Pokemon{}, err
		}

		return exploreJson, err
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return pokedex.Pokemon{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return pokedex.Pokemon{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return pokedex.Pokemon{}, err
	}

	c.cache.Add(url, body)

	exploreJson := pokedex.Pokemon{}
	err = json.Unmarshal(body, &exploreJson)
	if err != nil {
		return pokedex.Pokemon{}, err
	}

	return exploreJson, err
}
