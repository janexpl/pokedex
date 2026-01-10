package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocation(location string) (Locations, error) {
	var (
		data  []byte
		err   error
		found bool
	)
	url := baseURL + "location-area/" + location

	if data, found = c.cache.Get(url); !found {

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return Locations{}, err
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return Locations{}, err
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			return Locations{}, fmt.Errorf("received non-200 response code: %d", res.StatusCode)
		}
		data, err = io.ReadAll(res.Body)
		if err != nil {
			return Locations{}, fmt.Errorf("unable to read response body: %w", err)
		}
		c.cache.Add(url, data)
	}
	loc := Locations{}
	err = json.Unmarshal(data, &loc)
	if err != nil {
		return Locations{}, fmt.Errorf("unable to parse locations response: %w", err)
	}

	return loc, nil

}
