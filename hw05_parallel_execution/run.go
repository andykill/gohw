package hw05_parallel_execution //nolint:golint,stylecheck,revive

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func Run(tasks []Task, n int, m int) error {
	ch := make(chan Task, len(tasks))

	var errCnt int32
	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for {
				if atomic.LoadInt32(&errCnt) > int32(m) {
					return
				}
				task, ok := <-ch

				if !ok {
					return
				}

				if err := task(); err != nil {
					atomic.AddInt32(&errCnt, 1)
				}
			}
		}()
	}

	for _, t := range tasks {
		ch <- t
	}
	close(ch)

	wg.Wait()

	if errCnt > 0 && atomic.LoadInt32(&errCnt) >= int32(m) {
		return ErrErrorsLimitExceeded
	}

	return nil
}
