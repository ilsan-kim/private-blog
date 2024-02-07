package internal

import "time"

type PostMeta struct {
	ID          int
	Subject     string
	Preview     string
	Thumbnail   string
	FilePath    string
	CreatedTime time.Time
}

type PostService interface {
}
