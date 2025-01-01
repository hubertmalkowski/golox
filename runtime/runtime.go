package runtime

import "fmt"

type Runtime struct {
	hadError bool
}

func NewRuntime() *Runtime {
	return &Runtime{
		hadError: false,
	}
}

func (r *Runtime) Error(line int, message string) {
	r.report(line, "", message)
}

func (r *Runtime) report(line int, where string, message string) {
	fmt.Printf("[line %d] Error %s: %s", line, where, message)
	r.hadError = true
}
