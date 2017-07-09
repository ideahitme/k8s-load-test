package controller

import (
	"context"
	"fmt"
	"sync"

	"github.com/ideahitme/k8s-load-test/user"
)

var (
	defaultUsers      = 100
	defaultIterations = 10000
)

type Config struct {
	numUsers int
}

func Run(ctx context.Context, url string) error {
	done := make(chan struct{})
	fanIn := make(chan int)

	numIterations := defaultIterations
	if iters, ok := ctx.Value("iterations").(int); ok {
		numIterations = iters
	}

	numUsers := defaultUsers
	if users, ok := ctx.Value("users").(int); ok {
		numUsers = users
	}
	go func() {
		wg := sync.WaitGroup{}
		for i := 0; i < numUsers; i++ {
			wg.Add(1)
			go func() {
				for status := range user.New(done).Start(url) {
					fanIn <- status
				}
				wg.Done()
			}()
		}
		wg.Wait()
		close(fanIn)
	}()
	for status := range fanIn {
		fmt.Printf("received a status: %d\n", status)
		if numIterations == 0 {
			close(done)
		}
		numIterations--
		fmt.Printf("finished processing of %d statuses\n", (defaultIterations - numIterations))
	}
	return nil
}
