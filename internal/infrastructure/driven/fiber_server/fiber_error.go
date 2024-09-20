package fiber_server

import (
	"fmt"
	"service/internal/domain/server"
)

type ErrorDispatcher struct {
	errors []server.ValidationError
}

func NewErrorDispatcher() *ErrorDispatcher {
	return &ErrorDispatcher{}
}

func (e *ErrorDispatcher) AddError(err server.ValidationError) {
	e.errors = append(e.errors, err)
}

func (e *ErrorDispatcher) Error() string {
	total := len(e.errors)
	return fmt.Sprintf("There are %d errors", total)
}
