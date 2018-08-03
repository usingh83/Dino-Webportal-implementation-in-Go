package main

import (
	"dino/dinowebportal"
	"dino/databaselayer"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type configuration struct {
	Webserver string
}

func main() {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}
	config := new(configuration)
	json.NewDecoder(file).Decode(config)
	fmt.Println("Starting dino service on port ", config.Webserver)
	err = dinowebportal.RunWebPortal(databaselayer.MONGODB, config.Webserver, "mongodb://127.0.0.1", "dinowebportal/dinoTemplate")
	if err != nil {
		log.Panic(err)
	}
}
