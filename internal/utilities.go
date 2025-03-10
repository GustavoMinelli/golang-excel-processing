package internal

import (
	"encoding/json"
	"os"
)

// Configuration struct
type Configuration struct {
	Database struct {
		Host     string `json:"host"`
		User     string `json:"user"`
		Name     string `json:"name"`
		Password string `json:"password"`
		Sslmode  string `json:"ssl-mode"`
	} `json:"database"`
}

// GetConfig function
func GetConfig() (Configuration, error) {

	config := Configuration{}

	file, err := os.Open("./config.json")

	if err != nil {
		return config, err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)

	if err != nil {
		return config, err
	}

	return config, nil

}
