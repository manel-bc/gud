package fsutil

import (
	"errors"
	"os"
)

func Mkdir(dir string) error {
	err := os.Mkdir(dir, os.ModePerm)
	if errors.Is(err, os.ErrExist) {
		return nil
	}
	if err != nil {
		return err
	}

	return nil
}
