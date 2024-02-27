package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Config struct {
	DbConfig        DBConfig `json:"db_config"`
	FileWatcherPath string   `json:"file_watcher_path"`
}

type DBConfig struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	DB       string `json:"db"`
}

func MustLoadConfig(configPath string) Config {
	res := Config{}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Panicln(err)
	}

	f, err := os.ReadFile(configPath)
	if err != nil {
		log.Panicln(err)
	}

	err = json.Unmarshal(f, &res)
	if err != nil {
		log.Panicln(err)
	}

	return res
}

func (c DBConfig) ConnectionString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:5432/%s", c.User, c.Password, c.Host, c.DB)
}
