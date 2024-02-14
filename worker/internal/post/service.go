package post

import "github.com/ilsan-kim/private-blog/worker/internal/model"

type BaseService struct {
	repo Repository
}

func (b BaseService) Insert(data model.PostMeta) error {
	return b.repo.Insert(data)
}

func (b BaseService) Update(data model.PostMeta) error {
	orig, err := b.repo.Get(data.ID)
	if err != nil {
		return err
	}

	if data.Subject != "" {
		orig.Subject = data.Subject
	}

	if data.Preview != "" {
		orig.Preview = data.Preview
	}

	if data.Thumbnail != "" {
		orig.Thumbnail = data.Thumbnail
	}

	if data.FilePath != "" {
		orig.FilePath = data.FilePath
	}

	return b.repo.Update(orig)
}

func (b BaseService) Delete(id int) error {
	return b.repo.Delete(id)
}

func (b BaseService) Get(id int) (model.PostMeta, error) {
	return b.repo.Get(id)
}

func (b BaseService) GetAll() ([]model.PostMeta, error) {
	return b.repo.GetAll()
}
