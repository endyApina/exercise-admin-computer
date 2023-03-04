package main

import (
	"fmt"
	"log"

	"github.com/endyApina/exercise-admin-computer/config"
)

func main() {
	fmt.Println("running")
	config, err := config.LoadSecrets(".")
	if err != nil {
		log.Fatal("cannot load config")
	}

	log.Println(config.DatabaseName)
}
