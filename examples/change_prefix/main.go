package main

import (
	"encoding/json"

	"github.com/crgimenes/goConfig"
)

// Declare config struct

type mongoDB struct {
	Host string `cfgDefault:"example.com"`
	Port int    `cfgDefault:"999"`
}

type systemUser struct {
	Name     string `cfg:"name"`
	Password string `json:"passwd" cfg:"passwd"`
}

type configTest struct {
	Domain  string
	User    systemUser `cfg:"user"`
	MongoDB mongoDB
}

func main() {

	// Instance config struct
	config := configTest{}

	// Adds a prefix to the environment variables.
	goConfig.PrefixEnv = "EXAMPLE"

	// Pass the struct instance pointer to the parser
	err := goConfig.Parse(&config)
	if err != nil {
		println(err)
		return
	}

	// it just print the config struct on the screen
	j, _ := json.MarshalIndent(config, "", "\t")
	println(string(j))
}