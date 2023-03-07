package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type shell string

const (
	bash shell = "bash"
	zsh  shell = "zsh"
)

var allShells = []shell{
	bash,
	zsh,
}

func completion(cmd *cobra.Command, s shell) error {
	switch s {
	case bash:
		return cmd.GenBashCompletion(cmd.OutOrStdout())
	case zsh:
		return cmd.GenZshCompletion(cmd.OutOrStdout())
	}

	return fmt.Errorf("unsupported shell %q", s)
}

func addCompletionCommand(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use:   "completion",
		Short: "Generate shell completion",
	}
	rootCmd.AddCommand(cmd)

	add := func(s shell) {
		cmd.AddCommand(&cobra.Command{
			Use:   string(s),
			Short: "Generate " + string(s) + "completion",
			RunE: func(cmd *cobra.Command, args []string) error {
				return completion(rootCmd, s)
			},
		})
	}
	for _, s := range allShells {
		add(s)
	}
}
