package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func callbackMap(config *Config) error {
	res, err := http.Get("https://pokeapi.co/api/v2/location/")

	if err != nil {
		return fmt.Errorf("locations request error %v", err)
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		return fmt.Errorf("locations request body error %v", err)
	}

	var result []string

	err = json.Unmarshal(body, &result)

	if err != nil {
		return fmt.Errorf("unmarshal error %v", err)
	}

	return nil
}
