package post

import (
	"github.com/ilsan-kim/private-blog/worker/internal/model"
)

type Service interface {
	Insert(data model.PostMeta) error
	Update(filePath string, data model.PostMeta) error
	Delete(id int) error
	GetByFilePath(filePath string) (model.PostMeta, error)
	GetAll() ([]model.PostMeta, error)
}

type Repository interface {
	Insert(data model.PostMeta) error
	Update(id int, data model.PostMeta) error
	Delete(id int) error
	GetByFilePath(filePath string) (model.PostMeta, error)
	GetAll() ([]model.PostMeta, error)
}
