package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io"
)

type mapJson struct {
	Count int `json:"count"`
	Next *string `json:"next"`
	Previous *string `json:"previous"`
	Results []struct {
		Name string `json:"name"`
		Url string `json:"url"`
	} `json:"results"`
}

// URLs
// func getUrls() map[string]string {
// 	return map[string]string {
// 		"map": "https://pokeapi.co/api/v2/location-area/", 
// 	}
// }

func (c *Client) GetMap(pageURL *string) (mapJson, error) {
	// Create the ENDPOINT
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}
	
	var ret mapJson

	// returning if the cache has the value
	if val, ok := c.cache.Get(url); ok {
		err := json.Unmarshal(val, &ret)
		if err != nil {
			return mapJson{}, err
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