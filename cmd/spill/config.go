package main

import (
	"encoding/json"
	io "io/ioutil"
	"log"
)

type Config struct {
	Vault string
}

func ReadConfig() Config {
	var c Config

	config, readFileErr := io.ReadFile("config.json")
	if readFileErr != nil {
		log.Fatal(readFileErr)
	}

	jsonDecodeErr := json.Unmarshal(config, &c)
	if jsonDecodeErr != nil {
		log.Fatal(jsonDecodeErr)
	}

	return c
}
