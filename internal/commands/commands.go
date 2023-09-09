package commands

import (
	"context"
	"fmt"
	"github.com/manel-bc/gud/internal/models"
)

func Add(ctx context.Context, args []string) error {
	panic("not implemented")
}

func CatFile(ctx context.Context, args []string) error {
	panic("not implemented")
}

func Checkout(ctx context.Context, args []string) error {
	panic("not implemented")
}

func Commit(ctx context.Context, args []string) error {
	panic("not implemented")
}

func HashObject(ctx context.Context, args []string) error {
	panic("not implemented")
}

func Init(_ context.Context, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("init requires exactly one option but you provided: %v", args)
	}
	path := args[0]
	_, err := models.NewRepository(path, true)
	return err
}

func Log(ctx context.Context, args []string) error {
	panic("not implemented")
}

func LsTree(ctx context.Context, args []string) error {
	panic("not implemented")
}

func Merge(ctx context.Context, args []string) error {
	panic("not implemented")
}

func Rebase(ctx context.Context, args []string) error {
	panic("not implemented")
}

func RevParse(ctx context.Context, args []string) error {
	panic("not implemented")
}

func Rm(ctx context.Context, args []string) error {
	panic("not implemented")
}

func ShowRef(ctx context.Context, args []string) error {
	panic("not implemented")
}

func Tag(ctx context.Context, args []string) error {
	panic("not implemented")
}
