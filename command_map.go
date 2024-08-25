package main

import (
	"errors"
	"fmt"
	"log"
)

func callbackMap(cfg *config) error {

	res, err := cfg.pokeapiClient.ListLocalionAreas(cfg.nextLocationAreaURL)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Location areas:")

	for _, value := range res.Results {
		fmt.Println(value)
	}

	cfg.nextLocationAreaURL = res.Next
	cfg.prevLocationAreaURL = res.Previous

	return nil
}

func callbackMapb(cfg *config) error {

	if cfg.prevLocationAreaURL == nil {
		return errors.New("no previous page")
	}

	res, err := cfg.pokeapiClient.ListLocalionAreas(cfg.prevLocationAreaURL)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Location areas:")

	for _, value := range res.Results {
		fmt.Println(value)
	}

	cfg.nextLocationAreaURL = res.Next
	cfg.prevLocationAreaURL = res.Previous

	return nil
}
