package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := buildRootCommand().Execute(); err != nil {
		os.Exit(1)
	}
}

func buildRootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "scylla-cloud",
		Short: "Scylla Cloud server",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return initializeConfig(cmd)
		},
	}

	rootCmd.AddCommand(
		newAPIServerCommand(),
		newBookCommand(),
	)

	addCompletionCommand(rootCmd)

	return rootCmd
}

const envPrefix string = "SC"

// initializeConfig loops through all command flags and uses the appropriate
// environment value if no value is set by a command flag.
func initializeConfig(cmd *cobra.Command) error {
	v := viper.New()

	v.SetEnvPrefix(envPrefix)
	v.AutomaticEnv()

	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		if strings.Contains(f.Name, "-") {
			envVarSuffix := strings.ToUpper(strings.ReplaceAll(f.Name, "-", "_"))
			v.BindEnv(f.Name, fmt.Sprintf("%s_%s", envPrefix, envVarSuffix)) // nolint: errcheck
		}

		if !f.Changed && v.IsSet(f.Name) {
			val := v.Get(f.Name)
			cmd.Flags().Set(f.Name, fmt.Sprintf("%v", val)) // nolint: errcheck
		}
	})

	return nil
}
