package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	var (
		data  []byte
		found bool
		err   error
	)
	url := baseURL + "pokemon/" + name
	if data, found = c.cache.Get(url); !found {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return Pokemon{}, err
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return Pokemon{}, err
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			return Pokemon{}, fmt.Errorf("received non-200 response code: %d", res.StatusCode)
		}
		data, err = io.ReadAll(res.Body)
		if err != nil {
			return Pokemon{}, fmt.Errorf("unable to read response body: %w", err)
		}
		c.cache.Add(url, data)
	}
	poke := Pokemon{}
	err = json.Unmarshal(data, &poke)
	if err != nil {
		return Pokemon{}, fmt.Errorf("unable to parse pokemon response: %w", err)
	}

	return poke, nil
}
