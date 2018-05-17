package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path"

	"github.com/google/subcommands"
)

// sub is the interface for newCommand
type sub interface {
	SetFlags(f *flag.FlagSet)
	Execute(ctx context.Context, f *flag.FlagSet) subcommands.ExitStatus
}

type subcmd struct {
	sub
	name     string
	synopsis string
	usage    string
}

func (c subcmd) Name() string {
	return c.name
}

func (c subcmd) Synopsis() string {
	return c.synopsis
}

func (c subcmd) Usage() string {
	return c.usage + "\nFlags:\n"
}

func (c subcmd) Execute(ctx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	return c.sub.Execute(ctx, f)
}

// handleError returns subcommands.ExitSuccess if err is nil.
// If err is non-nil, it returns subcommands.ExitFailure.
func handleError(err error) subcommands.ExitStatus {
	if err == nil {
		return subcommands.ExitSuccess
	}
	fmt.Fprintf(os.Stderr, "\nError: %v\n", err)
	return subcommands.ExitFailure
}

// newCommander creates a subcommands.Commander for nested sub commands.
// This registers "flags" and "help" sub commands for the new commander.
func newCommander(f *flag.FlagSet, name string) *subcommands.Commander {
	name = fmt.Sprintf("%s %s", path.Base(os.Args[0]), name)
	c := subcommands.NewCommander(f, name)
	c.Register(c.FlagsCommand(), "misc")
	c.Register(c.HelpCommand(), "misc")
	return c
}
