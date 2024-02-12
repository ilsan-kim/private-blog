package watcher

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

var ErrNotFile = errors.New("not a file")

type FileWatcher struct {
	path        string
	prev        []DiffItem
	diffHandler DiffHandler
}

type File struct {
	Name        string
	CreatedTime time.Time
	Content     string
}

func NewFileWatcher(path string) FileWatcher {
	// TODO path 외부 주입
	return FileWatcher{path: path, prev: nil, diffHandler: PostDiffHandler{}}
}

func (f FileWatcher) Watch(stop context.CancelFunc) Watcher {
	// TODO : sleep threshold 컨피그로 빼기
	defer time.Sleep(2 * time.Second)

	var err error
	infos, err := f.getDiffItems()
	log.Printf("prev > %d, now > %d\n", len(f.prev), len(infos))
	if err != nil {
		stop()
		return f
	}

	if len(f.prev) == 0 {
		// 첫 검사
		f.prev = infos
		return f
	} else {
		// 첫 검사 이후
		diffs := f.diffHandler.FindDiff(f.prev, infos)
		if len(diffs) > 0 {
			log.Println("diff found!!!!!!!!!!!!")
			log.Println(diffs)
			stop()
			return f
		}

		f.prev, err = f.getDiffItems()
		if err != nil {
			// TODO stop 할때 기본 정보를 백업해야함 (prev 정보)
			stop()
			return f
		}
	}

	return f
}

func (f FileWatcher) getDiffItem(filePath string) (DiffItem, error) {
	path := fmt.Sprintf("%s/%s", f.path, filePath)

	// TODO: diffMode 의 상수화
	diffItem, err := NewDiffItem("CONTENT", path)
	return diffItem, err
}

func (f FileWatcher) listDirectory() ([]os.DirEntry, error) {
	entries, err := os.ReadDir(f.path)
	if err != nil {
		return nil, err
	}

	return entries, nil
}

func (f FileWatcher) getDiffItems() ([]DiffItem, error) {
	res := make([]DiffItem, 0)

	files, err := f.listDirectory()
	if err != nil {
		log.Printf("can not list files in directory %s\n", f.path)
		return nil, err
	}

	for _, file := range files {
		postDiffItem, err := f.getDiffItem(file.Name())
		if err != nil {
			log.Printf("can not initialize FileContentDiffItem for %s\n", file.Name())
			continue
		}

		res = append(res, postDiffItem)
	}

	return res, nil
}

func (f FileWatcher) handleDiff() error {
	now, err := f.getDiffItems()
	if err != nil {
		return err
	}

	f.diffHandler.FindDiff(f.prev, now)
	return nil
}

func (f FileWatcher) gracefullyStop(stop context.CancelFunc) {
	// TODO: implement
}
