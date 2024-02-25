package main

import (
	"fmt"
	"log"
	"os"

	"github.com/leesolway/powerwave/src/config"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading configuration:", err)
	}

	if len(os.Args) > 1 {
		ExecuteCLI()
	} else {
		router := SetupRouter()
		router.Run(fmt.Sprintf(":%d", config.Port))
	}
}
