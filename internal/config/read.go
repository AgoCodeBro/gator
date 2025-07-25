package config

import(
	"os"
	"fmt"
	"encoding/json"
)

func Read() (Config, error) {
	configFile, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}
	
	data, err := os.ReadFile(configFile)
	if err != nil {
		return Config{}, err
	}

	var result Config
	if err := json.Unmarshal(data, &result); err != nil {
		return Config{}, err
	}

	return result, nil
}


func getConfigFilePath() (string, error) {
	userHome, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	file_path := fmt.Sprintf("%v/.gatorconfig.json", userHome)

	return file_path, nil
}
