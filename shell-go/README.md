# Codecrafters Shell in Go

This project is a simple implementation of a shell written in Go, built as part of the Codecrafters challenge. It demonstrates command parsing, tokenization, and execution of built-in shell commands like `cd`, `pwd`, `echo`, and more.

---

## Features

- **Built-in Commands**:
   - `cd`: Change the current directory.
   - `pwd`: Print the current working directory.
   - `echo`: Print text to the shell output.
   - `exit`: Exit the shell.
- **Custom Tokenizer**:
   - Parses user input into tokens for command execution.
- **Command Manager**:
   - Manages built-in commands and their execution.
- **Interactive Shell**:
   - Reads user input from standard input and executes commands.

---

## Project Structure

```plaintext
codecrafters-shell-go/
├── cmd/
│   └── myshell/
│       └── main.go               # Entry point of the shell
├── pkg/
│   └── command/
│       ├── handlers/
│       │   ├── cd.go             # Implementation of the 'cd' command
│       │   ├── echo.go           # Implementation of the 'echo' command
│       │   ├── exit.go           # Implementation of the 'exit' command
│       │   ├── pwd.go            # Implementation of the 'pwd' command
│       ├── command.go            # Core command logic
│       ├── command-manager.go    # Manages command execution
│       ├── registry.go           # Registry for built-in commands
├── reader/
│   └── reader.go                 # Reads user input
├── tokenizer/
│   └── parser.go                 # Tokenizes user input
└── .gitattributes                # Git attributes configuration