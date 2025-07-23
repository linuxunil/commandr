package commandr

import "context"

// BaseCall provides a basic implementation of the Call interface.
// It can be embedded in custom call types to extend functionality.
type BaseCall struct {
	Name    string
	Args    []string
	Context context.Context
}

func NewCall(name string, args []string, ctx context.Context) *BaseCall {
	return &BaseCall{
		Name:    name,
		Args:    args,
		Context: ctx,
	}
}

// Name returns the command name.
func (c *BaseCall) GetName() string { return c.Name }

// Args returns the command arguments.
func (c *BaseCall) GetArgs() []string { return c.Args }

// Context returns the command context.
func (c *BaseCall) GetContext() context.Context { return c.Context }
