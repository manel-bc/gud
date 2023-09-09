package commands

import (
	"github.com/manel-bc/gud/internal/models"
)

type Init struct {
	Path string `arg:"" help:"Create an empty Git repository" type:"path"`
}

func (i *Init) Run(_ *Context) error {
	_, err := models.NewRepository(i.Path, true)
	return err
}
