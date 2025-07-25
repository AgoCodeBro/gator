package config

import (
	"os"
	"encoding/json"
)

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
