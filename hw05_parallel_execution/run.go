package hw05_parallel_execution //nolint:golint,stylecheck

import (
	"errors"
	"log"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func Run(tasks []Task, n int, m int) error {
	ch := make(chan Task)
	producer(tasks, ch)

	mu := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(n)
	var errCnt int
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			for {
				task, ok := <-ch
				if !ok {
					log.Println("Chanel close")
					break
				}

				err := task()
				if err != nil {
					mu.Lock()
					errCnt++
					mu.Unlock()
				}
				if errCnt >= m {
					break
				}
			}
		}()
	}

	wg.Wait()

	if errCnt > 0 && errCnt >= m {
		return ErrErrorsLimitExceeded
	}
	return nil
}
func producer(tasks []Task, ch chan<- Task) {
	for _, t := range tasks {
		ch <- t
	}
	close(ch)
}
