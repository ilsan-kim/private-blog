package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	MDFileUploadPath    string   `json:"md_file_upload_path"`
	ThumbnailUploadPath string   `json:"thumbnail_upload_path"`
	ImageUploadFrom     []string `json:"image_upload_from"`
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
