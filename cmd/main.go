package main

import (
	"log"

	"github.com/prasanth-pn/cleancode/pkg/config"
	di "github.com/prasanth-pn/cleancode/pkg/dependencies"
)

func main() {

	// load config
	cnf, err := config.LoadConfig()
	if err != nil {
		log.Fatal("failed to load environements")
	}

	// server initialization
	server, err := di.InitiallizeEvent(cnf)
	if err != nil {
		log.Fatal("failed to initialize the files")
	}

	
	server.Start(cnf)
}
