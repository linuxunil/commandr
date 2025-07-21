package commandr

import (
	"fmt"
	"strings"
)

// Clean input for processing of commands.
func tokenizeInput(text string) []string {
	fields := strings.Fields(text)
	return fields
}

// Stores information about the command
type commandMetaData struct {
	name        string
	description string
}

// Commands with a callback and optional args.
type Command[T any] struct {
	meta     commandMetaData
	args     []string
	callback func(conf *T, opt ...[]string) error
}

// Registry of Commands available
// Add new Commands with the register function
type Commands[T any] struct {
	registered map[string]func(*T, Command[T]) error
}

func (c *Commands[T]) Add(name string, f func(*T, Command[T]) error) {
	c.registered[name] = f
}

// Look up and execute a command.
func (c *Commands[T]) Execute(st *T, cmd Command[T]) error {
	com, ok := c.registered[cmd.meta.name]
	if !ok {
		return fmt.Errorf("Command %v does not exist", cmd.meta.name)
	}
	if err := com(st, cmd); err != nil {
		return err
	}
	return nil
}

func New[T any]() *Commands[T] {
	return &Commands[T]{
		registered: make(map[string]func(*T, Command[T]) error),
	}
}

// func (c *Commander[T]) Register(name string, handler func(*T, []string) error) {
// 	// Adapter to convert between handler signatures
// 	c.register(name, func(state *T, cmd command) error {
// 		return handler(state, cmd.args)
// 	})
// }
//
// func (c *Commander[T]) Execute(state *T, input string) error {
// 	tokens := CleanInput(input)
// 	if len(tokens) == 0 {
// 		return fmt.Errorf("no command provided")
// 	}
//
// 	cmd := command{
// 		name: tokens[0],
// 		args: tokens[1:],
// 	}
//
// 	return c.run(state, cmd)
// }

