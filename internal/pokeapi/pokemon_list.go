package pokeapi

func (c *Client) ListPokemons(location *string) (RespShallowLocations, error) {