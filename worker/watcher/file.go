package watcher

import (
	"context"
	"database/sql"
	"errors"
	"github.com/ilsan-kim/private-blog/worker/internal/model"
	"github.com/ilsan-kim/private-blog/worker/internal/post"
	"github.com/ilsan-kim/private-blog/worker/pkg"
	"log"
	"time"
)

type FileHandler interface {
	handle(meta model.PostMeta) error
}

func NewFileHandler(postService post.Service, mark int) (FileHandler, error) {
	switch mark {
	case pkg.DiffResultMarkAdded:
		return FileAddHandler{postService}, nil
	case pkg.DiffResultMarkEdited:
		return FileEditHandler{postService}, nil
	case pkg.DiffResultMarkDeleted:
		return FileDeleteHandler{postService}, nil
	default:
		return nil, errors.New("undefined file handler")
	}
}

type FileAddHandler struct{ service post.Service }
type FileEditHandler struct{ service post.Service }
type FileDeleteHandler struct{ service post.Service }

func (h FileAddHandler) handle(post model.PostMeta) error {
	log.Printf("handle add post %v\n", post)
	return h.service.Insert(post)
}

func (h FileEditHandler) handle(post model.PostMeta) error {
	log.Printf("handle edit post %v\n", post)
	return h.service.Update(post.FilePath, post)
}

func (h FileDeleteHandler) handle(post model.PostMeta) error {
	log.Printf("handle delete post %v\n", post)
	return h.service.Delete(post.ID)
}

type FileWatcher struct {
	path        string
	diffHandler pkg.FileDiffHandler
	postService post.Service
}

func NewFileWatcher(path string, postRepo post.Repository) FileWatcher {

	// TODO path 외부 주입
	return FileWatcher{
		path:        path,
		diffHandler: pkg.FileDiffHandler{DiffMode: "CONTENT", DirPath: path},
		postService: post.NewBaseService(postRepo),
	}
}

func (f FileWatcher) Watch(stop context.CancelFunc) Watcher {
	// TODO : sleep threshold 컨피그로 빼기
	defer time.Sleep(2 * time.Second)

	var err error

	f.diffHandler, err = f.diffHandler.HandleDiff(func(diffs []pkg.DiffResult) error {
		for _, d := range diffs {
			file := model.NewFile(f.path, d.Item.GetName(), d.Item.GetTime())
			p, err := f.postService.GetByFilePath(file.Name)
			if errors.Is(err, sql.ErrNoRows) {
				err = nil
				p = model.PostMetaFromFile(file, "")
			}
			if err != nil {
				log.Println(err)
				continue
			}

			handler, err := NewFileHandler(f.postService, d.Mark)
			if err != nil {
				log.Println(err)
				continue
			}

			err = handler.handle(p)
			if err != nil {
				log.Println(err)
				continue
			}
		}
		return nil
	})

	if err != nil {
		log.Println(err)
		stop()
	}

	return f
}
