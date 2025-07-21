package commandr

import (
	"context"
	"errors"
	"strings"
)

type BaseCommand struct {
	Name        string
	Description string
	Aliases     []string
}

type Command[T any] interface {
	Run(ctx context.Context, state *T, args []string) (string, error)
}

// Registery of Commands available
// Add new Commands with the register function
type Commands[T any] struct {
	commands map[string]Command[T]
}

var (
	ErrNoCommand   = errors.New("no command provided")
	ErrNotFound    = errors.New("command does not exist")
	ErrInvalidArgs = errors.New("invalid arguments")
)

// Return a new register
func New[T any]() *Commands[T] {
	return &Commands[T]{
		commands: make(map[string]Command[T]),
	}
}

// Register a Command with the registery.
func (c *Commands[T]) Add(cmd Command[T]) {
	cmdName := strings.ToLower(cmd.Name())
	c.commands[cmdName] = cmd
	for _, alias := range cmd.Aliases() {
		c.commands[strings.ToLower(alias)] = cmd
	}

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

// Clean input for processing of Commands.
func tokenizeInput(text string) []string {
	fields := strings.Fields(text)
	return fields
}

// func (c *Commander[T]) Register(name string, handler func(*T, []string) error) {
// 	// Adapter to convert between handler signatures
// 	c.register(name, func(state *T, cmd Command) error {
// 		return handler(state, cmd.args)
// 	})
// }
//
// func (c *Commander[T]) Execute(state *T, input string) error {
// 	tokens := CleanInput(input)
// 	if len(tokens) == 0 {
// 		return fmt.Errorf("no Command provided")
// 	}
//
// 	cmd := Command{
// 		name: tokens[0],
// 		args: tokens[1:],
// 	}
//
// 	return c.run(state, cmd)
// }

