package stackerr

import (
	"fmt"
	"strings"
)

// Errors ...
type Errors struct {
	List         []error
	FormatFunc   func(string, []error) string
	FormatForRow string
}

// DefaultFormatter is a default formatter func
func DefaultFormatter(rowfmt string, e []error) string {
	rows := []string{}
	for i, err := range e {
		rows = append(rows, fmt.Sprintf(rowfmt, i, err))
	}
	return strings.Join(rows, "\n")
}

// New initialize Errors.
func New(errs ...error) *Errors {
	return &Errors{
		List: errs,
	}
}

// Err returns nil if no any error pooled.
func (e *Errors) Err() error {
	if len(e.List) == 0 {
		return nil
	}
	return e
}

// IfExists is an alias for Err().
func (e *Errors) IfExists() error {
	return e.Err()
}

// Push pushes other errors to the list.
func (e *Errors) Push(errs ...error) *Errors {
	e.List = append(e.List, errs...)
	return e
}

// Pushf pushes another error with formatting.
func (e *Errors) Pushf(format string, a ...interface{}) *Errors {
	return e.Push(fmt.Errorf(format, a...))
}

// Error to implement builtin.error interface.
func (e *Errors) Error() string {
	if e.FormatFunc == nil {
		e.FormatFunc = DefaultFormatter
	}
	if e.FormatForRow == "" {
		e.FormatForRow = "[%d]\t%v"
	}
	return e.FormatFunc(e.FormatForRow, e.List)
}
