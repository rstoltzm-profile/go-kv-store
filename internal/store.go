package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func PutKeyValue(fullFilePath, key, value string) {
	// Read existing content
	kvMap := make(map[string]string)

	if _, err := os.Stat(fullFilePath); err == nil {
		file, err := os.Open(fullFilePath)
		check(err)
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			parts := strings.SplitN(line, "=", 2)
			if len(parts) == 2 {
				kvMap[parts[0]] = parts[1]
			}
		}
	}

	// Update or add the key
	kvMap[key] = value

	// Write back to file
	file, err := os.Create(fullFilePath)
	check(err)
	defer file.Close()

	for k, v := range kvMap {
		_, err := file.WriteString(fmt.Sprintf("%s=%s\n", k, v))
		check(err)
	}
}

func GetKeyValue(fullFilePath, key string) (string, bool) {
	file, err := os.Open(fullFilePath)
	if err != nil {
		return "", false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 && parts[0] == key {
			return parts[1], true
		}
	}

	return "", false
}
