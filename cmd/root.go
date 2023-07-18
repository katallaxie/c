package cmd

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/katallaxie/g/internal/cfg"

	"github.com/katallaxie/pkg/env"
	"github.com/spf13/cobra"
)

var config = cfg.Default()

func init() {
	RootCmd.AddCommand(InitCmd)

	RootCmd.PersistentFlags().BoolVarP(&config.Verbose, "verbose", "v", config.Verbose, "verbose output")
	RootCmd.PersistentFlags().StringVarP(&config.Template, "template", "t", config.Template, "template to use")
	RootCmd.PersistentFlags().BoolVarP(&config.Force, "force", "f", config.Force, "force overwrite")

	RootCmd.SilenceErrors = true
	RootCmd.SilenceUsage = true
}

var RootCmd = &cobra.Command{
	Use:   "g",
	Short: "g is a tiny scaffolding tool",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runRoot(cmd.Context(), args...)
	},
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return preRunRoot(cmd.Context(), args...)
	},
}

func preRunRoot(ctx context.Context, args ...string) error {
	cwd, err := config.Cwd()
	if err != nil {
		return err
	}

	if len(args) == 0 {
		return fmt.Errorf("no folder name provided")
	}

	p := filepath.Clean(filepath.Join(cwd, args[0]))

	checker := env.NewChecker()
	err = checker.Check(ctx, env.IsDirEmpty(p))
	if err != nil {
		return err
	}

	return nil
}

func runRoot(ctx context.Context, args ...string) error {
	return nil
}
