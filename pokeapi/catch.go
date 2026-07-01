package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io"
)

type Pokemon struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Base_experience int `json:"base_experience"`

	Height int `json:"height"`
	Weight int `json:"weight"`

	Stats []struct{
		BaseStat int `json:"base_stat"`
		Stat struct{
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`

	Types []struct{
		Type struct{
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}

// URLs
// func getUrls() map[string]string {
// 	return map[string]string {
// 		"map": "https://pokeapi.co/api/v2/location-area/", 
// 	}
// }

func (c *Client) GetPokemon(poke string) (Pokemon, error) {
	// Create the ENDPOINT
	url := baseURL + "/pokemon/" + poke 
	
	var ret Pokemon

	// returning if the cache has the value
	if val, ok := c.cache.Get(url); ok {
		err := json.Unmarshal(val, &ret)
		if err != nil {
			return Pokemon{}, err
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