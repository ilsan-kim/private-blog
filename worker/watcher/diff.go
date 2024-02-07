package watcher

import "github.com/ilsan-kim/private-blog/worker/internal"

type DiffItem interface{
	Diff() bool
}

type DiffHandler interface {
	Find() []DiffItem
	Handle(prev, now []DiffItem) error
}

type PostDiffItem struct {
	hash []byte
}

func (i PostDiffItem) Diff() bool  {
	return true
}

type PostDiffHandler struct {
	postService internal.PostService
}

func (h PostDiffHandler) Find() []DiffItem {
	return []DiffItem{}
}

func (h PostDiffHandler) Handle(prev, now []DiffItem) error {
	return nil
}
