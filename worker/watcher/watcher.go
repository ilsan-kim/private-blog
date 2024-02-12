package watcher

import "context"

type Watcher interface {
	Watch(ctx context.CancelFunc) Watcher
	// HandleDiff() error
}
