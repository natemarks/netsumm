package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Target is a struct that contains information about a target to be polled
type Target struct {
	Type        string `json:"type"`
	Destination string `json:"destination"`
	Data        string `json:"data"` // data for a target that's specific to the target type
}

// Config is a struct that contains information about the targets to be polled
type Config struct {
	Targets    []Target `json:"targets"`
	Source     string   `json:"source"`
	Iterations int      `json:"iterations"`
}

func parseConfig(jsonString string) (Config, error) {
	var config Config

	err := json.Unmarshal([]byte(jsonString), &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}

func readFileContents(filePath string) (string, error) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Read the file content
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	// Convert file content to string
	fileContentStr := string(content)
	return fileContentStr, nil
}

// GetConfigFromFile returns a Config struct
func GetConfigFromFile() Config {
	fileContents, err := readFileContents("netsumm.json")
	if err != nil {
		panic(err)
	}
	config, err := parseConfig(fileContents)
	if err != nil {
		panic(err)
	}
	return config
}
