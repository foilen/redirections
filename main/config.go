package main

import "encoding/json"

// Configuration is the json configuration file
type Configuration struct {
	Host        string
	Redirection string
	Permanent   bool
	AppendQuery bool
}

func getConfiguration(jsonText string) (*Configuration, error) {
	var configuration Configuration
	jsonBytes := []byte(jsonText)
	err := json.Unmarshal(jsonBytes, &configuration)

	return &configuration, err
}
