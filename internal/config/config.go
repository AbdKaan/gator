package config

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
)

const configFileName string = "/.gatorconfig.json"

type Config struct {
	Db_url            string `json:"db_url"`
	Current_user_name string `json:"current_user_name"`
}

func (c *Config) SetUser(name string) error {
	c.Current_user_name = name
	return write(c)
}

func Read() (Config, error) {
	home_dir, err := os.UserHomeDir()
	if err != nil {
		return Config{}, err
	}

	file_dir := filepath.Join(home_dir, configFileName)
	file, err := os.Open(file_dir)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return Config{}, err
	}

	var config Config
	if err := json.Unmarshal(byteValue, &config); err != nil {
		return Config{}, err
	}

	return config, nil
}

func write(c *Config) error {
	byte_config, err := json.Marshal(c)
	if err != nil {
		return err
	}

	home_dir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	file_dir := filepath.Join(home_dir, configFileName)
	err = os.WriteFile(file_dir, byte_config, 0644)
	if err != nil {
		return err
	}

	return nil
}
