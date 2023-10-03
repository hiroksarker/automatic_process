package app

import (
	"fmt"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"project_root/app/error_handler"
)

type ServerManager struct {
	ErrorHandler error_handler.ErrorHandler
}

func NewServerManager() *ServerManager {
	return &ServerManager{
		ErrorHandler: error_handler.NewErrorHandler(),
	}
}

func (sm *ServerManager) LoadAndExecute() error {
	configData, err := ioutil.ReadFile("config/server_config.yml")
	if err != nil {
		return sm.ErrorHandler.HandleError(fmt.Sprintf("Config file 'config/server_config.yml' not found: %v", err))
	}

	config := make(map[string]interface{})
	if err := yaml.Unmarshal(configData, &config); err != nil {
		return sm.ErrorHandler.HandleError(fmt.Sprintf("Error parsing YAML file: %v", err))
	}

	servers := config["servers"]
	tasks := config["tasks"]

	// Handle each server and tasks here

	return nil
}