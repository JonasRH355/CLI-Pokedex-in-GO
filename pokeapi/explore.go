package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io"
)

type Pokemons struct {
	Pokemon_encounters []struct{
		Pokemon pokemon `json:"pokemon"`
		
	} `json:"pokemon_encounters"`

}

type pokemon struct {
	Name string `json:"name"`
	Url string `json:"url"`
}

// https://pokeapi.co/api/v2/location-area/{id or name}/
// https://pokeapi.co/api/v2/location-area/canalave-city-area?field=PokemonEncounter

func (c *Client) GetExploreRegion(name string) (Pokemons, error) {
	// Create the ENDPOINT
	url := baseURL + "/location-area/" + name
	
	var ret Pokemons

	// returning if the cache has the value
	if val, ok := c.cache.Get(url); ok {
		err := json.Unmarshal(val, &ret)
		if err != nil {
			return Pokemons{}, err
		}

		return ret, nil
	}
	
	// Create the request
	req, err := http.NewRequest("GET",url,nil)
	if err != nil {
		return ret,err
	}

	//Do the request
	res, err := c.httpClient.Do(req)
	if err != nil {
		fmt.Println(err, " [GET]")
		return ret, err
	}
	defer res.Body.Close()

	// Read the request
	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return ret, err
	}

	// Unmarshal the request with the MapJson
	err = json.Unmarshal(dat, &ret)
	if err != nil {
		return ret, err
	}

	c.cache.Add(url, dat)
	return ret, nil
}