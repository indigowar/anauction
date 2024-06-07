package models

import "strings"

type AggregatedError struct {
	Errors []error
}

var _ error = &AggregatedError{}

// Error implements error.
func (e *AggregatedError) Error() string {
	builder := strings.Builder{}
	for _, err := range e.Errors {
		builder.WriteString(err.Error())
		builder.WriteString("\n")
	}

	return builder.String()
}
