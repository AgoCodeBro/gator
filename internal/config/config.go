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

func SetUser(name string) error  {
	c, err := Read()
	if err != nil {
		return err
	}

	c.CurrentUserName = name

	if err := write(c); err != nil {
		return err
	}

	return nil
}

func write(cfg Config) error {
	jsonData, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	
	file, err := getConfigFilePath()
	if err != nil {
		return err
	}

	err = os.WriteFile(file, jsonData, 0644)
	if err != nil {
		return err
	}
	
	return nil
}
