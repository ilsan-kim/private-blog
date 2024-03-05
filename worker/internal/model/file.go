package model

import (
	"time"
)

type File struct {
	Name        string
	CreatedTime time.Time
	Content     string
}

func NewFile(fileName string, createdTime time.Time) File {
	return File{
		Name:        fileName,
		CreatedTime: createdTime,
	}
}
