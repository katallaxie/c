package cmd

import (
	"context"

	"github.com/spf13/cobra"
)

func init() {

}

var RootCmd = &cobra.Command{
	Use:   "g",
	Short: "g is a tiny scaffolding tool",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runRoot(cmd.Context())
	},
}

func runRoot(ctx context.Context) error {
	return nil
}
