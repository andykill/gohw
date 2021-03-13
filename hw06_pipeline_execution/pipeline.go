package hw06_pipeline_execution //nolint:golint,stylecheck,revive

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) (out Out) {
	ch := make(chan Stage)
	defer close(ch)

	go func() {
		select {
		case <-done:
			return
		default:
		}
		for {
			stage, ok := <-ch
			if !ok {
				break
			}
			in = stage(in)
		}
	}()

	for _, stage := range stages {
		ch <- stage
	}
	out = in
	return out
}
