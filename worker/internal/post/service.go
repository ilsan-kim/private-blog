package post

import "github.com/ilsan-kim/private-blog/worker/internal/model"

type BaseService struct {
	repo Repository
}

func NewBaseService(repo Repository) BaseService {
	return BaseService{repo: repo}
}

func (b BaseService) Insert(data model.PostMeta) error {
	return b.repo.Insert(data)
}

func (b BaseService) Update(subject string, data model.PostMeta) error {
	post, err := b.repo.GetBySubject(subject)
	if err != nil {
		return err
	}

	return b.repo.Update(post.ID, data)
}

func (b BaseService) Delete(id int) error {
	return b.repo.Delete(id)
}

func (b BaseService) GetByFilePath(filePath string) (model.PostMeta, error) {
	return b.repo.GetBySubject(filePath)
}

func (b BaseService) GetAll() ([]model.PostMeta, error) {
	return b.repo.GetAll()
}
