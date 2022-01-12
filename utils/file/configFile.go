package file

import (
	"bufio"
	"io"
	"os"
	"strings"
)

type Config map[string]string

func ReadConfigTXT(filename string) (Config, error) {
	config := Config{
		"user":       "barca",
		"password":   "abc123",
		"port":       "8080",
		"ip":         "127.0.0.1",
		"database":   "barcaDB",
		"dbname":     "test-db",
		"frammework": "gin",
		"language":   "portugues-brasileiro",
	}
	if len(filename) == 0 {
		return config, nil
	}

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		// check if the line to file '.env'
		if equal := strings.Index(line, "="); equal >= 0 {
			if key := strings.TrimSpace(line[:equal]); len(key) > 0 {
				value := ""
				if len(line) > equal {
					value = strings.TrimSpace(line[equal+1:])
				}
				// assign the config map
				config[key] = value
			}
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
	}
	return config, nil
}
