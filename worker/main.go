package main

import (
	"context"
	watcher2 "github.com/ilsan-kim/private-blog/worker/watcher"
	"log"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	worker := 2

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	var wg sync.WaitGroup

	var w watcher2.Watcher
	w = watcher2.NewFileWatcher("md/posts")

	for i := 0; i < worker; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			work(ctx, stop, w)
		}()
	}

	<-ctx.Done()

	log.Println("got interruption signal")
	wg.Wait()
}

func work(ctx context.Context, stop context.CancelFunc, w watcher2.Watcher) {
	for {
		select {
		case <-ctx.Done():
			log.Println("work done")
			return
		default:
			err := w.Watch(stop)
			if err != nil {
				log.Println(err)
				return
			}
		}
	}
}
