package post

type BaseService struct {
	repo Repository
}

func (b BaseService) Insert(data PostMeta) error {
	return b.repo.Insert(data)
}

func (b BaseService) Update(data PostMeta) error {
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

func (b BaseService) Get(id int) (PostMeta, error) {
	return b.repo.Get(id)
}

func (b BaseService) GetAll() ([]PostMeta, error) {
	return b.repo.GetAll()
}
