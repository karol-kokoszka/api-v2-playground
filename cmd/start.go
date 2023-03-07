package cmd

import (
	"github.com/spf13/cobra"

	"github.com/scylladb/scylla-cloud/internal/restapi"
)

type startCmd struct {
	cobra.Command
}

func newStartCommand() *cobra.Command {
	cmd := startCmd{
		Command: cobra.Command{
			Use:   "start",
			Short: "starts external API service",
		},
	}
	cmd.RunE = func(c *cobra.Command, args []string) error {
		return cmd.createAndStartHTTPServer()
	}
	return &cmd.Command
}

func (cmd *startCmd) createAndStartHTTPServer() (err error) {
	router := restapi.NewRouter()
	router.Run(":8080")
	return
}
