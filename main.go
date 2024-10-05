package main

import (
	"fmt"
	"log"
	"os"

	"github.com/AbdKaan/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error trying to read config: %v", err)
	}

	config_state := state{&cfg}

	var handlerFunctions = make(map[string]func(*state, command) error)
	cli_commands := commands{handlerFunctions}

	cli_commands.register("login", handlerLogin)

	arguments := os.Args
	if len(arguments) < 2 {
		fmt.Println("not enough arguments given")
		os.Exit(1)
	}

	cmd := command{arguments[1], arguments[2:]}
	err = cli_commands.run(&config_state, cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
