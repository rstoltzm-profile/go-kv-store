package internal

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Cluster string `json:"cluster"`
}

func LoadConfig(configPath string) Config {
	data, err := os.ReadFile(configPath)
	check(err)

	var config Config
	err = json.Unmarshal(data, &config)
	check(err)

	return config
}

func ExecuteAction(action, fullFilePath, key, value string) {
	switch action {
	case "put":
		if value == "" {
			fmt.Println("Error: value is required for put action")
			os.Exit(1)
		}
		PutKeyValue(fullFilePath, key, value)
		fmt.Printf("Stored: %s = %s\n", key, value)
	case "get":
		val, found := GetKeyValue(fullFilePath, key)
		if found {
			fmt.Printf("%s = %s\n", key, val)
		} else {
			fmt.Printf("Key '%s' not found\n", key)
		}
	default:
		fmt.Printf("Error: unknown action '%s'. Use 'get' or 'put'\n", action)
		os.Exit(1)
	}
}
