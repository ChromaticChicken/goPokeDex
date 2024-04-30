package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (mapsJson, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if body, ok := c.cache.Get(url); ok {
		locationsResp := mapsJson{}
		err := json.Unmarshal(body, &locationsResp)
		if err != nil {
			return mapsJson{}, err
		}

		return locationsResp, err
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return mapsJson{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return mapsJson{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return mapsJson{}, err
	}

	c.cache.Add(url, body)

	locationsResp := mapsJson{}
	err = json.Unmarshal(body, &locationsResp)
	if err != nil {
		return mapsJson{}, err
	}

	return locationsResp, err
}
