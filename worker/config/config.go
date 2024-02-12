package config

import "fmt"

type DBConfig struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	DB       string `json:"db"`
}

func (c DBConfig) ConnectionString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:5432/%s", c.User, c.Password, c.Host, c.DB)
}
