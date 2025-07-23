package commandr

import (
	"errors"
	"strings"
)

// Common errors returned by the command system.
var (
	// ErrNoCommand is returned when no command is provided.
	ErrNoCommand = errors.New("no command provided")

	// ErrNotFound is returned when the requested command does not exist.
	ErrNotFound = errors.New("command does not exist")

	// ErrInvalidArgs is returned when command arguments are invalid.
	ErrInvalidArgs = errors.New("invalid arguments")
)

// DefaultCommands is the default command registry used by package-level functions.
var DefaultCommands = &Commands{}

type CommandFunc func(res Result, req Call)

// Command represents a registered command with its pattern and handler.
type Command struct {
	header  Header
	pattern string
	handler CommandFunc
}

// Commands is a command registry that manages command registration and execution.
// It provides thread-safe registration and lookup of commands.
type Commands struct {
	commands map[string]Command
}

func (f CommandFunc) Exec(res Result, req Call) {
	f(res, req)
}

// Exec executes the command with the provided context, result, and call.
// It returns an error if execution fails.
func Exec(res Result, req Call) {
	DefaultCommands.Exec(res, req)
}

func (c *Commands) Exec(res Result, req Call) {
	cmd, _ := c.findCommand(req)
	cmd.Exec(res, req)
}

// HandleFunc registers a command handler with the given pattern.
// The handler function will be called when the command is executed.
func (c *Commands) HandleFunc(pattern string, handler func(res Result, req Call)) {
	if c.commands == nil {
		c.commands = make(map[string]Command)
	}
	cmd := Command{
		header:  &BaseHeader{},
		pattern: pattern,
		handler: CommandFunc(handler),
	}
	c.commands[pattern] = cmd
}

func HandleFunc(pattern string, handler func(res Result, req Call)) {
	if DefaultCommands.commands == nil {
		DefaultCommands.commands = make(map[string]Command)
	}
	cmd := Command{
		header:  &BaseHeader{},
		pattern: pattern,
		handler: CommandFunc(handler),
	}
	DefaultCommands.commands[pattern] = cmd
}

func (c *Commands) findCommand(req Call) (CommandFunc, error) {
	n := req.GetName()
	if cf, exists := c.commands[n]; exists {
		return cf.handler, nil
	}
	return nil, ErrNotFound
}

// tokenizeInput splits input text into command tokens.
// // It removes extra whitespace and returns a slice of string tokens.
// func tokenizeInput(text string) []string {
// 	fields := strings.Fields(text)
// 	return fields
// }
