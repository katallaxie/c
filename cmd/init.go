package cmd

import (
	"context"
	"os"
	"path/filepath"

	"github.com/katallaxie/g/pkg/spec"

	"github.com/katallaxie/pkg/utils/files"
	"github.com/spf13/cobra"
)

var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new config",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runInit(cmd.Context())
	},
}

func runInit(_ context.Context) error {
	if err := spec.Write(spec.Default(), spec.DefaultFilename, config.Force); err != nil {
		return err
	}

	cwd, err := config.Cwd()
	if err != nil {
		return err
	}

	path := filepath.Clean(filepath.Join(cwd, spec.DefaultFilename))

	err = files.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
