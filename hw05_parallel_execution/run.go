package hw05_parallel_execution //nolint:golint,stylecheck,revive

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func Run(tasks []Task, n int, m int) error {
	ch := make(chan Task, len(tasks))

	var errCnt int
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for {
				mu.Lock()
				if errCnt > m {
					mu.Unlock()
					return
				}
				mu.Unlock() // хз как подругому (((
				task, ok := <-ch

				if !ok {
					return
				}

				if err := task(); err != nil {
					mu.Lock()
					errCnt++
					mu.Unlock()
				}
			}
		}()
	}

	for _, t := range tasks {
		ch <- t
	}
	close(ch)

	wg.Wait()

	if errCnt > 0 && errCnt >= m {
		return ErrErrorsLimitExceeded
	}

	return nil
}
