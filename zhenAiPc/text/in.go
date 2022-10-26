package text

import "fmt"

type ConcurrentEngine struct {
	Scheduler Scheduler
}
type Request struct {
	Url string
}

type Scheduler interface {
	Run()
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	for _, r := range seeds {
		fmt.Printf(r.Url)
	}
	const t = string("Run")
	//return t
}
