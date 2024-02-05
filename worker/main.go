package main

import (
	"context"
	"log"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	worker := 2

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			work(ctx)
		}()
	}

	<-ctx.Done()

	log.Println("got interruption signal")
	wg.Wait()
}

func work(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			log.Println("work done")
			return
		default:
			time.Sleep(4 * time.Second)
			log.Println("im now working")
		}
	}
}
