package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"kv-store/internal"
)

func main() {
	filename := flag.String("filename", "kvstore.txt", "the file to store key-value pairs")
	key := flag.String("key", "", "the key")
	value := flag.String("value", "", "the value")
	action := flag.String("action", "get", "action to perform: get or put")
	server := flag.String("server", "", "server path")
	configPath := flag.String("config", "configs/servers.json", "path to config file")

	flag.Parse()

	if *key == "" {
		fmt.Println("Error: key is required")
		os.Exit(1)
	}

	config := internal.LoadConfig(*configPath)
	fullFilePath := filepath.Join(config.Cluster, *server, *filename)

	internal.ExecuteAction(*action, fullFilePath, *key, *value)
}
