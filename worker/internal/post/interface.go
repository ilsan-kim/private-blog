package post

import (
	"github.com/ilsan-kim/private-blog/worker/internal/model"
)

type Service interface {
	Insert(data model.PostMeta) error
	Update(data model.PostMeta) error
	Delete(id int) error
	Get(id int) (model.PostMeta, error)
	GetAll() ([]model.PostMeta, error)
}

type Repository interface {
	Insert(data model.PostMeta) error
	Update(data model.PostMeta) error
	Delete(id int) error
	Get(id int) (model.PostMeta, error)
	GetAll() ([]model.PostMeta, error)
}
