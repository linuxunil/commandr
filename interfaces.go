// Package commandr provides a lightweight, generic command parsing and execution
// library for Go applications.
package commandr

import (
	"context"
)

// Handler represents a command handler that can execute commands.
// Implementations should process the call and populate the result.
type Handler interface {
	// Exec executes the command with the given result and call context.
	// It should populate the result with output, error, and duration information.
	Exec(r Result, c Call)
}

type Header interface {
	Del(k string) bool
	Set(k string, v []byte)
	Get(k string) ([]byte, error)
	Has(k string) bool
}

// Call represents a command invocation with its name, arguments, and context.
// Implementations can extend this interface to include additional metadata.
type Call interface {
	// Name returns the command name.
	GetName() string

	// Args returns the command arguments as a slice of strings.
	GetArgs() []string

	// Context returns the context for the command execution,
	// which can be used for cancellation, timeouts, and request-scoped values.
	GetContext() context.Context
}

// Result represents the outcome of a command execution.
// It provides methods to get and set output, error, and timing information.
type Result interface {
	// GetOutput returns the command's output string.
	GetOutput() string

	// SetOutput sets the command's output string.
	SetOutput(out string)
}
