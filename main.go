package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/hiroksarker/automatic_process/pkg/sync"
)

func main() {
	// Initialize and start the file watcher
	watcher := sync.NewFileWatcher("config/server_config.yaml")
	go watcher.Watch()

	// Create a signal channel to handle interrupts
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)

	// Start the server manager
	serverManager := app.NewServerManager()

	// Periodically check for updated configurations
	for {
		select {
		case <-watcher.Changes:
			// Configuration file has changed
			fmt.Println("Configuration file changed. Reloading...")
			if err := serverManager.ReloadConfiguration(); err != nil {
				fmt.Printf("Error reloading configuration: %v\n", err)
			}
		case <-signalChannel:
			// Handle interrupt signal
			fmt.Println("Received interrupt signal. Shutting down...")
			serverManager.Shutdown()
			return
		}
	}
}
