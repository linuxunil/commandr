# Commandr

A lightweight, generic command parsing and execution library for Go applications.

## Overview

Commandr provides a simple, type-safe way to register and execute commands in CLI applications. Built with Go generics, it works with any application state type while maintaining clean separation between command parsing and business logic.

## Features

- **Generic Design** - Works with any state type using Go generics
- **Simple API** - Easy command registration and execution
- **Error Handling** - Commands return errors for proper error propagation
- **Input Parsing** - Built-in input tokenization and cleaning
- **Zero Dependencies** - Uses only Go standard library

## Quick Start

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/linuxunil/commandr"
)

// Define your application state
type AppState struct {
    username string
    loggedIn bool
}

// Create command handlers
func loginHandler(state *AppState, args []string) error {
    if len(args) < 1 {
        return fmt.Errorf("username required")
    }
    state.username = args[0]
    state.loggedIn = true
    fmt.Printf("Logged in as %s\n", state.username)
    return nil
}

func statusHandler(state *AppState, args []string) error {
    if state.loggedIn {
        fmt.Printf("Logged in as %s\n", state.username)
    } else {
        fmt.Println("Not logged in")
    }
    return nil
}

func main() {
    // Create state and commander
    state := &AppState{}
    cmdr := commandr.New[AppState]()
    
    // Register commands
    cmdr.Register("login", loginHandler)
    cmdr.Register("status", statusHandler)
    
    // Execute commands
    if err := cmdr.Execute(state, "login alice"); err != nil {
        log.Fatal(err)
    }
    
    if err := cmdr.Execute(state, "status"); err != nil {
        log.Fatal(err)
    }
}
```

## API Reference

### Creating a Commander

```go
cmdr := commandr.New[YourStateType]()
```

### Registering Commands

```go
cmdr.Register(name string, handler func(*YourStateType, []string) error)
```

### Executing Commands

```go
err := cmdr.Execute(state *YourStateType, input string) error
```

## Input Format

Commands are parsed from strings with the following format:
- First word is the command name (case-insensitive)
- Remaining words are arguments
- Whitespace is automatically trimmed

Examples:
- `"login alice"` → command: "login", args: ["alice"]
- `"move north quickly"` → command: "move", args: ["north", "quickly"]
- `"  LOOK  "` → command: "look", args: []

## Error Handling

Commands should return errors for invalid input or execution failures:

```go
func myHandler(state *MyState, args []string) error {
    if len(args) < 2 {
        return fmt.Errorf("expected 2 arguments, got %d", len(args))
    }
    // ... command logic
    return nil
}
```

## Development

### Requirements

- Go 1.24.4 or later
- mise (optional, for task management)

### Running Tests

```bash
# With mise
mise run test

# Or directly
go test ./...
```

### Running with Coverage

```bash
mise run test-cover
# or
go test -cover ./...
```

### Formatting and Linting

```bash
mise run check
# or manually:
go fmt ./...
go vet ./...
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Add tests for new functionality
4. Ensure all tests pass
5. Submit a pull request

## License

MIT License - see LICENSE file for details

## Design Philosophy

Commandr was extracted from common patterns observed in multiple Go CLI applications. It prioritizes:

- **Simplicity** - Minimal API surface
- **Type Safety** - Leverages Go generics for compile-time safety
- **Reusability** - Works with any application state
- **Testability** - Easy to unit test command handlers
- **Performance** - Zero-allocation command lookup

## Inspiration

This library was inspired by command patterns from:
- CLI argument parsing libraries
- MUD (Multi-User Dungeon) command systems
- Interactive shell implementations

The goal is to provide the command handling portion without the complexity of full CLI frameworks.