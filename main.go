package main

import (
	"fmt"
	"github.com/go-yaml/yaml"
	"log"
	"project_root/app"
)

func main() {
	serverManager := app.NewServerManager()
	if err := serverManager.LoadAndExecute(); err != nil {
		log.Fatalf("An unexpected error occurred: %v", err)
	}
}