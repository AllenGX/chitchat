package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
	Address      string
	ReadTimeout  int64
	WriteTimeout int64
	Static       string
}

func LoadConfig(config *Configuration) {
	file, err := os.Open("config.json")
	if err != nil {
		fmt.Println("Cannot open config file", err.Error())
		panic(err)
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(config)
	if err != nil {
		fmt.Println("Cannot get configuration from file", err.Error())
		panic(err)
	}
}
