package main

import (
	"context"
	"github.com/ilsan-kim/private-blog/worker/config"
	"github.com/ilsan-kim/private-blog/worker/internal/post"
	watcher2 "github.com/ilsan-kim/private-blog/worker/watcher"
	"log"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	config.UseJsonLogger()

	// TODO: config
	dbConf := config.DBConfig{
		Host:     "localhost",
		User:     "postgres",
		Password: "postgres",
		DB:       "blog_dev",
	}

	worker := 1

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	var wg sync.WaitGroup
	var w watcher2.Watcher

	postRepo := post.NewPGRepository(dbConf)

	w = watcher2.NewFileWatcher("md/posts", postRepo)

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
			w = w.Watch(stop)
		}
	}
}
