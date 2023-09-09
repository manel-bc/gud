package models

import (
	"errors"
	"fmt"
	"github.com/go-ini/ini"
	"github.com/manel-bc/gud/internal/fsutil"
	"os"
	"path/filepath"
)

var (
	gitDirName     = ".git"
	configFileName = "config"

	repoVersionKeyName = "repositoryformatversion"
)

// Repository represents a git repository
type Repository struct {
	workTree string
	gitDir   string
	config   *ini.File
}

func NewRepository(path string, create bool) (Repository, error) {
	repo := Repository{
		workTree: path,
		gitDir:   filepath.Join(path, gitDirName),
		config:   ini.Empty(),
	}

	if create {
		if err := repo.Create(); err != nil {
			return Repository{}, fmt.Errorf("failed to create git repository: %v", err)
		}
	}

	cfg, err := ini.Load(filepath.Join(repo.gitDir, configFileName))
	if err != nil {
		return Repository{}, fmt.Errorf("failed to load config: %v", err)
	}
	repo.config = cfg

	ver, err := repo.config.Section("core").GetKey(repoVersionKeyName)
	if err != nil {
		return Repository{}, fmt.Errorf("failed to read config: %v", err)
	}
	if ver.Value() != "0" {
		return Repository{}, fmt.Errorf("unsuported %s: %s", repoVersionKeyName, ver.Value())
	}

	return repo, nil
}

func (r Repository) Create() error {
	if err := r.createWorktree(r.workTree); err != nil {
		return err
	}

	if err := fsutil.Mkdir(r.gitDir); err != nil {
		return fmt.Errorf("failed to create git directory: %v", err)
	}

	for _, d := range []string{
		"branches",
		"objects",
		"refs",
		"refs/tags",
		"refs/heads",
	} {
		path := filepath.Join(r.gitDir, d)
		if err := fsutil.Mkdir(path); err != nil {
			return err
		}
	}

	descriptionPath := filepath.Join(r.gitDir, "description")
	descriptionContent := []byte("Unnamed repository; edit this file 'description' to name the repository.\n")
	if err := os.WriteFile(descriptionPath, descriptionContent, 0644); err != nil {
		return fmt.Errorf("failed to write description file: %v", err)
	}

	headPath := filepath.Join(r.gitDir, "HEAD")
	headContent := []byte("ref: refs/heads/master\n")
	if err := os.WriteFile(headPath, headContent, 0644); err != nil {
		return fmt.Errorf("failed to write description file: %v", err)
	}

	if err := r.createConfigFile(); err != nil {
		return fmt.Errorf("failed to create config file: %v", err)
	}

	return nil
}

func (Repository) createWorktree(path string) error {
	if err := os.Mkdir(path, os.ModePerm); err != nil && !errors.Is(err, os.ErrExist) {
		return fmt.Errorf("failed to create worktree: %v", err)
	}
	ls, err := os.ReadDir(path)
	if err != nil {
		return err
	}
	if len(ls) != 0 {
		return errors.New("directory is not empty")
	}

	return nil
}

func (r Repository) createConfigFile() error {
	cfg := ini.Empty()

	core, err := cfg.NewSection("core")
	if err != nil {
		return fmt.Errorf("failed to create new config section: %v", err)
	}

	for k, v := range map[string]string{
		"repositoryformatversion": "0",
		"filemode":                "false",
		"bare":                    "false",
	} {
		if _, err = core.NewKey(k, v); err != nil {
			return fmt.Errorf("failed to add config section '%s': %v", k, err)
		}
	}

	path := filepath.Join(r.gitDir, configFileName)
	return cfg.SaveToIndent(path, "\t")
}
