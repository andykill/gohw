package hw06_pipeline_execution //nolint:golint,stylecheck,revive

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) (out Out) {
	for _, stage := range stages {
		pereliv := make(chan interface{})

		go func(in In, pereliv chan interface{}) {
			defer close(pereliv)

			for {
				select {
				case <-done:
					return
				case val, ok := <-in:
					if !ok {
						return
					}
					pereliv <- val
				}
			}
		}(in, pereliv)

		in = stage(pereliv)
	}

	return in
}
