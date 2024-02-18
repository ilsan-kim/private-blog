package model

import (
	"fmt"
	"time"
)

type File struct {
	Name        string
	CreatedTime time.Time
	Content     string
}

func NewFile(dirPath, fileName string, createdTime time.Time) File {
	return File{
		Name:        fmt.Sprintf("%s/%s", dirPath, fileName),
		CreatedTime: time.Now(),
	}
}
