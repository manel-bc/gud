package main

import (
	"context"
	"flag"
	"github.com/manel-bc/gud/internal/commands"
	"log"
	"os"

	"github.com/peterbourgon/ff/v3/ffcli"
)

func main() {
	root := &ffcli.Command{
		Name:       "gud",
		ShortUsage: "gud <SUBCOMMAND> [OPTIONS]",
		FlagSet:    flag.NewFlagSet("gud", flag.ExitOnError),
		Subcommands: []*ffcli.Command{
			{
				Name: "add",
				Exec: commands.Add,
			},
			{
				Name: "cat-file",
				Exec: commands.CatFile,
			},
			{
				Name: "checkout",
				Exec: commands.Checkout,
			},
			{
				Name: "commit",
				Exec: commands.Commit,
			},
			{
				Name: "hash-object",
				Exec: commands.HashObject,
			},
			{
				Name:       "init",
				Exec:       commands.Init,
				ShortUsage: "init [path]",
			},
			{
				Name: "log",
				Exec: commands.Log,
			},
			{
				Name: "ls-tree",
				Exec: commands.LsTree,
			},
			{
				Name: "merge",
				Exec: commands.Merge,
			},
			{
				Name: "rebase",
				Exec: commands.Rebase,
			},
			{
				Name: "rev-parse",
				Exec: commands.RevParse,
			},
			{
				Name: "rm",
				Exec: commands.Rm,
			},
			{
				Name: "show-ref",
				Exec: commands.ShowRef,
			},
			{
				Name: "tag",
				Exec: commands.Tag,
			},
		},
	}

	if err := root.ParseAndRun(context.Background(), os.Args[1:]); err != nil {
		log.Fatalf("git gud: %v", err)
	}
}
