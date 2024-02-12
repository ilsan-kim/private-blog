package post

import "time"

type PostMeta struct {
	ID          int
	Subject     string
	Preview     string
	Thumbnail   string
	FilePath    string
	CreatedTime time.Time
}

type Service interface {
}

type Repository interface {
	Insert(data PostMeta) error
	Update(data PostMeta) error
	Delete(pkey int) error
	Get(pkey int) PostMeta
	GetAll() []PostMeta
}
