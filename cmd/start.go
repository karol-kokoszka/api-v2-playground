package cmd

import (
	"net/http"

	"github.com/spf13/cobra"

	"github.com/scylladb/scylla-cloud/api"
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
	cmd.Flags().String("router", "chi", "router to use")
	cmd.RunE = func(c *cobra.Command, args []string) error {
		switch c.Flag("router").Value.String() {
		case "chi":
			return cmd.createAndStartChiHTTPServer()
		default:
			return cmd.createAndStartHTTPServer()
		}
	}
	return &cmd.Command
}

func (cmd *startCmd) createAndStartHTTPServer() (err error) {
	router := restapi.NewRouter()
	router.Run(":8080")
	return
}

func (cmd *startCmd) createAndStartChiHTTPServer() (err error) {
	api.RegisterRoutes(http.DefaultServeMux)
	return http.ListenAndServe(":8080", nil)
}
