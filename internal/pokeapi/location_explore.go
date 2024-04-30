package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ExploreLocation(name string) (locationJson, error) {
	url := baseURL + "/location-area/" + name

	if body, ok := c.cache.Get(url); ok {
		exploreJson := locationJson{}
		err := json.Unmarshal(body, &exploreJson)
		if err != nil {
			return locationJson{}, err
		}

		return exploreJson, err
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return locationJson{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return locationJson{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return locationJson{}, err
	}

	c.cache.Add(url, body)

	exploreJson := locationJson{}
	err = json.Unmarshal(body, &exploreJson)
	if err != nil {
		return locationJson{}, err
	}

	return exploreJson, err
}
