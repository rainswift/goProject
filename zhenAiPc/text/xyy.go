package text

import "fmt"

type Structs struct {
	Fun     Fun
	TextXyy string
}

type Fun interface {
	XyyFn()
}

func (e *Structs) XyyFn(val string) {
	fmt.Printf(val + "\n")
	fmt.Printf(e.TextXyy)
}
