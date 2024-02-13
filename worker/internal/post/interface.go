package post

import "time"

type PostMeta struct {
	ID          int
	Subject     string
	Preview     string
	Thumbnail   string
	FilePath    string
	CreatedTime time.Time
	UpdatedTime time.Time
}

type Service interface {
	Insert(data PostMeta) error
	Update(data PostMeta) error
	Delete(id int) error
	Get(id int) (PostMeta, error)
	GetAll() ([]PostMeta, error)
}

type Repository interface {
	Insert(data PostMeta) error
	Update(data PostMeta) error
	Delete(id int) error
	Get(id int) (PostMeta, error)
	GetAll() ([]PostMeta, error)
}
