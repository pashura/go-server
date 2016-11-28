package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
	Port string
	Name string
}

func GetConfig() Configuration {
	file, _ := os.Open("./conf.json")
	fmt.Println(file)
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(configuration.Port)

	return configuration
}

func main() {
	GetConfig()
}
