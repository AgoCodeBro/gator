package config

import ("os"
				"fmt"
				"encoding/json"
				)

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}


func Read() (Config, error) {
	configFilePath, err := getConfigPath()
	if err != nil {
		return Config{}, err
	}

	fileData, err := os.ReadFile(configFilePath)
	if err != nil {
		return Config{}, err
	}

	var config Config
	if err := json.Unmarshal(fileData, &config); err != nil {
		return Config{}, err
	}

	return config, nil
}


func (c Config) SetUser(userName string) (error) {
	c.CurrentUserName = userName
	err := write(c)
	return err
}

func write(cnfg Config) (error) {
	jsonData, err := json.Marshal(cnfg)
	if err != nil {
		return err
	}

	configFilePath, err := getConfigPath()
	if err != nil {
		return err
	}
	
	if err := os.WriteFile(configFilePath, jsonData, 0666); err != nil {
		return err
	}

	return nil

}

func getConfigPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	configFilePath := fmt.Sprintf("%v/.gatorconfig.json", homeDir)
	
	return configFilePath, nil
}

