package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	var (
		data  []byte
		err   error
		found bool
	)
	url := baseURL + "location-area/"
	if pageURL != nil {
		url = *pageURL
	}
	if data, found = c.cache.Get(url); !found {

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return RespShallowLocations{}, err
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return RespShallowLocations{}, err
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			return RespShallowLocations{}, fmt.Errorf("received non-200 response code: %d", res.StatusCode)
		}
		data, err = io.ReadAll(res.Body)
		if err != nil {
			return RespShallowLocations{}, fmt.Errorf("unable to read response body: %w", err)
		}
		c.cache.Add(url, data)
	}
	locations := RespShallowLocations{}
	err = json.Unmarshal(data, &locations)
	if err != nil {
		return RespShallowLocations{}, fmt.Errorf("unable to parse locations response: %w", err)
	}

	return locations, nil

}
