package config

import(
	"os"
	"fmt"
	"encoding/json"
)

type Config struct {
	DbURL  string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	configFile, err := getConfigFilePath()
	if err != nil {
		return Config{}, fmt.Errorf("Failed to get the config file path: %v", err)
	}
	
	data, err := os.ReadFile(configFile)
	if err != nil {
		if os.IsNotExist(err) {
			// Create file
			f, err := os.Create(configFile)
			if err != nil {
				return Config{}, fmt.Errorf("Config file not found and failed to create one")
			}

			// Create an empty json object in the file as a placeholder for when it gets unmarshaled.
			_, err = f.WriteString("{}")
			if err != nil {
				return Config{}, fmt.Errorf("Config file not found and failed to add empty json to it")
			}

		} else {
			return Config{}, fmt.Errorf("Failed to read the config file: %v", err)
		}
	}

	var result Config
	if err := json.Unmarshal(data, &result); err != nil {
		return Config{}, fmt.Errorf("Failed to unmarshal the config json: %v", err)
	}

	return result, nil
}


func getConfigFilePath() (string, error) {
	userHome, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("Failed to find user home dir: %v", err)
	}

	file_path := fmt.Sprintf("%v/.gatorconfig.json", userHome)

	return file_path, nil
}

func SetUser(name string) error  {
	c, err := Read()
	if err != nil {
		return fmt.Errorf("Failed to read config file: %v", err)
	}

	c.CurrentUserName = name

	if err := write(c); err != nil {
		return fmt.Errorf("Failed to write to config file: %v", err)
	}

	return nil
}

func write(cfg Config) error {
	jsonData, err := json.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("Failed to marshal json: %v", err)
	}
	
	file, err := getConfigFilePath()
	if err != nil {
		return fmt.Errorf("Failed to find config file: %v", err)
	}

	err = os.WriteFile(file, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("Failed to write to %v: %v", file, err)
	}
	
	return nil
}
