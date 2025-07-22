package commandr

import (
	"context"
	"errors"
	"strings"
)

var (
	ErrNoCommand   = errors.New("no command provided")
	ErrNotFound    = errors.New("command does not exist")
	ErrInvalidArgs = errors.New("invalid arguments")
)
var DefaultCommands = New[any]()

type Runner[T any] interface {
	Run(ctx context.Context, state *T, args []string) (string, error)
}

type BaseCommand struct {
	Name        string
	Description string
	Aliases     []string
}

type Command[T any] struct {
	meta    BaseCommand
	handler func(ctx context.Context, state *T, args []string) (string, error)
}

// Registery of Commands available
// Add new Commands with the register function
type Commands[T any] struct {
	commands map[string]Command[T]
}

// Return a new register
func New[T any]() *Commands[T] {
	return &Commands[T]{
		commands: make(map[string]Command[T]),
	}
}

// Register a Command with the registery.
func (c *Commands[T]) Add(cmd Command[T]) {
	cmdName := strings.ToLower(cmd.meta.Name)
	c.commands[cmdName] = cmd
	for _, alias := range cmd.meta.Aliases {
		c.commands[strings.ToLower(alias)] = cmd
	}
}

func (cmd Command[T]) Run(ctx context.Context, state *T, args []string) (string, error) {
	return cmd.handler(ctx, state, args)
}

// Look up and execute a Command.
func (reg *Commands[T]) Execute(ctx context.Context, state *T, input string) (string, error) {
	tokens := tokenizeInput(input)
	if len(tokens) == 0 {
		return "", ErrNoCommand
	}

	cmdToken, args := strings.ToLower(tokens[0]), tokens[1:]

	cmd, ok := reg.commands[cmdToken]
	if !ok {
		return "", ErrNotFound
	}
	return cmd.Run(ctx, state, args)
}

func HandleFunc(name, description string, handler func(ctx context.Context, state any, args []string) (string, error), aliases ...string) {
	cmd := Command[any]{
		meta: BaseCommand{
			Name:        name,
			Description: description,
			Aliases:     aliases,
		},
		handler: handler,
	}
	DefaultCommands.Add(cmd)
}

// Clean input for processing of Commands.
func tokenizeInput(text string) []string {
	fields := strings.Fields(text)
	return fields
}
