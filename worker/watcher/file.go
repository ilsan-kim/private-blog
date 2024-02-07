package watcher

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"time"
)

var ErrNotFile = errors.New("not a file")

type FileWatcher struct {
	path        string
	prev        []File
	diffHandler DiffHandler
}

type File struct {
	Name        string
	CreatedTime time.Time
	Content     string
}

func NewFileWatcher(path string) FileWatcher {
	// TODO path, postService 외부 주입
	return FileWatcher{path: path, prev: nil}
}

func (f FileWatcher) Watch(stop context.CancelFunc) error {
	files, err := f.listDirectory()
	if err != nil {
		// TODO stop 할때 기본 정보를 백업해야함 (prev 정보)
		stop()
		return err
	}

	// TODO 이거 컨피그로
	time.Sleep(2 * time.Second)
	log.Println(len(files), " watching")
	return nil
}

func (f FileWatcher) handleDiff() {
	// TODO: implement
}

func (f FileWatcher) gracefullyStop(stop context.CancelFunc) {
	// TODO: implement
}

func (f FileWatcher) getFileMetas() ([]File, error) {
	var res []File

	entries, err := f.listDirectory()
	if err != nil {
		return nil, err
	}

	for _, e := range entries {
		f, err := f.getFileMeta(e)
		if err != nil {
			return nil, err
		}
		// for debugging
		log.Println(f.Name)
		log.Println(f.Content[0:int(math.Min(30, float64(len(f.Content))))])
		log.Println(f.CreatedTime)

		log.Println("=====")

		res = append(res, f)
	}

	return res, nil
}

func (f FileWatcher) listDirectory() ([]os.DirEntry, error) {
	entries, err := os.ReadDir(f.path)
	if err != nil {
		return nil, err
	}

	return entries, nil
}

func (f FileWatcher) getFileMeta(file os.DirEntry) (File, error) {
	res := File{}
	if file.IsDir() {
		// Nested Directory ..
		return res, ErrNotFile
	}

	res.Name = file.Name()

	info, err := file.Info()
	if err != nil {
		return res, err
	}
	res.CreatedTime = info.ModTime()

	content, err := os.ReadFile(fmt.Sprintf("%s/%s", f.path, res.Name))
	if err != nil {
		return res, err
	}
	res.Content = string(content)

	return res, nil
}
