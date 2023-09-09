package main

import (
	"github.com/alecthomas/kong"
	"github.com/manel-bc/gud/internal/commands"
)

var cli struct {
	Init commands.Init `cmd:"" help:"Initializes a git repository"`
}

func main() {
	ctx := kong.Parse(&cli)
	pathToRepo := ""
	err := ctx.Run(&commands.Context{RepositoryRoot: pathToRepo})
	ctx.FatalIfErrorf(err)
}
