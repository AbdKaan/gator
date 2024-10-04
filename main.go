package main

import (
	"fmt"
	"log"

	"github.com/AbdKaan/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error trying to read config: %v", err)
	}

	cfg.SetUser("canko")

	cfg2, err := config.Read()
	if err != nil {
		log.Fatalf("error trying to read config: %v", err)
	}
	fmt.Println(cfg2)
}
