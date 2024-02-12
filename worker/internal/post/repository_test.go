package post

import (
	"github.com/ilsan-kim/private-blog/worker/config"
	"log"
	"testing"
)

func TestNewPGRepository(t *testing.T) {
	conf := config.DBConfig{
		Host:     "localhost",
		User:     "postgres",
		Password: "postgres",
		DB:       "postgres",
	}

	repo := NewPGRepository(conf)
	err := repo.db.Ping()
	log.Println(err)
}
