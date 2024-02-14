package watcher

import (
	"context"
	"github.com/ilsan-kim/private-blog/worker/internal/model"
	"github.com/ilsan-kim/private-blog/worker/internal/post"
	"github.com/ilsan-kim/private-blog/worker/pkg"
	"log"
	"time"
)

type FileWatcher struct {
	path        string
	prev        []pkg.DiffItem
	diffHandler pkg.FileDiffHandler
	postService post.Service
}

func NewFileWatcher(path string) FileWatcher {

	// TODO path 외부 주입
	return FileWatcher{
		path:        path,
		prev:        nil,
		diffHandler: pkg.FileDiffHandler{DiffMode: "CONTENT", DirPath: path},
	}
}

func (f FileWatcher) Watch(stop context.CancelFunc) Watcher {
	// TODO : sleep threshold 컨피그로 빼기
	defer time.Sleep(2 * time.Second)

	var err error

	f.diffHandler, err = f.diffHandler.HandleDiff(func(diffs []pkg.DiffResult) error {
		for _, d := range diffs {
			file := model.NewFile(f.path, d.Item.GetName())
			p := model.PostMetaFromFile(file, "")
			if d.Mark == pkg.DiffResultMarkAdded {
				log.Printf("handle added file... %v\n", p)
			} else if d.Mark == pkg.DiffResultMarkDeleted {
				log.Printf("handle deleted file...%v\n", p)
			} else if d.Mark == pkg.DiffResultMarkEdited {
				p.UpdatedTime = time.Now()
				log.Printf("handle updated file...%v\n", p)
			}

			log.Printf("%d: %s", d.Mark, d.Item.GetName())
		}
		return nil
	})

	if err != nil {
		log.Println(err)
		stop()
	}

	return f
}
