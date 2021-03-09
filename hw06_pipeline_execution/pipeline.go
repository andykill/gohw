package hw06_pipeline_execution //nolint:golint,stylecheck,revive
import "sync"

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) (out Out) {
	ch := make(chan Stage)
	defer close(ch)

	var wg sync.WaitGroup
	wg.Add(len(stages))

	go func() {
		defer wg.Done()
		for {
			select {
			case <-done:
				return
			case stage, ok := <-ch:
				if !ok {
					return
				}
				in = stage(in)
			}
		}
	}()

	for _, stage := range stages {
		ch <- stage
	}
	wg.Wait()
	out = in
	return out
}
