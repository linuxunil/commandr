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

// Commands with a callback and optional args.
type command[T any] struct {
	name        string
	description string
	callback    func(conf *T, opt ...[]string) error
}

// Registery of Commands available
// Add new Commands with the register function
type Commands[T any] struct {
	registered map[string]command[T]
}

// Register a command with the registery.
func (c *Commands[T]) Add(name string, callback func(conf *T, opt ...[]string) error, description ...string) {
	cmd := command[T]{name: name, callback: callback}

	if len(description) > 0 {
		cmd.description = strings.Join(description, " ")
	}
	c.registered[strings.ToLower(name)] = cmd
}

// Look up and execute a command.
func (reg *Commands[T]) Execute(st *T, input string) error {
	tokens := tokenizeInput(input)
	if len(tokens) == 0 {
		return fmt.Errorf("No command provided")
	}

	cmdToken, args := strings.ToLower(tokens[0]), tokens[1:]

	cmd, ok := reg.registered[cmdToken]
	if !ok {
		return fmt.Errorf("Command %v does not exist", cmdToken)
	}
	if err := cmd.callback(st, args); err != nil {
		return err
	}
	return nil
}

func New[T any]() *Commands[T] {
	return &Commands[T]{
		registered: make(map[string]command[T]),
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
