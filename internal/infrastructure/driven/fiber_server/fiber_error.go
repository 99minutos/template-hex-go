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

func (a *ErrorDispatcher) AddError(err server.ValidationError) {
	a.errors = append(a.errors, err)
}

func (a *ErrorDispatcher) Error() string {
	total := len(a.errors)
	return fmt.Sprintf("There are %d errors", total)
}
