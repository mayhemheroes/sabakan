package main

import (
	"context"
	"flag"
	"net/http"
	"os"

	"github.com/cybozu-go/cmd"
	"github.com/cybozu-go/sabakan/client"
	"github.com/google/subcommands"
)

var (
	flagServer = flag.String("server", "http://localhost:8888", "<Listen IP>:<Port number>")
)

func main() {
	c := client.NewClient(*flagServer, &cmd.HTTPClient{
		Client: &http.Client{},
	})

	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(ipamCommand(c), "")
	subcommands.Register(machinesCommand(c), "")

	flag.Parse()
	cmd.LogConfig{}.Apply()

	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}
