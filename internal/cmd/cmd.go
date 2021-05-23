package cmd

import (
	"io"
	"os"
)

// BaseCommand provides common fields to all commands.
type BaseCommand struct {
	Stdin  io.ReadCloser
	Stdout io.WriteCloser
	Stderr io.WriteCloser
}

// Init initializes BaseCommand with default arguments
func (b *BaseCommand) Init() {
	b.Stdin = os.Stdin
	b.Stdout = os.Stdout
	b.Stderr = os.Stderr
}