package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListPokemons(location string) (LocationPokemons, error) {
	var (
		data  []byte
		err   error
		found bool
	)
	url := baseURL + "location-area/" + location

	if data, found = c.cache.Get(url); !found {

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return LocationPokemons{}, err
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return LocationPokemons{}, err
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			return LocationPokemons{}, fmt.Errorf("received non-200 response code: %d", res.StatusCode)
		}
		data, err = io.ReadAll(res.Body)
		if err != nil {
			return LocationPokemons{}, fmt.Errorf("unable to read response body: %w", err)
		}
		c.cache.Add(url, data)
	}
	pokemons := LocationPokemons{}
	err = json.Unmarshal(data, &pokemons)
	if err != nil {
		return LocationPokemons{}, fmt.Errorf("unable to parse locations response: %w", err)
	}

	return pokemons, nil

}
