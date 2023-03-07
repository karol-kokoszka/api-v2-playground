package cmd

import (
	pb "github.com/scylladb/scylla-cloud/gen/proto/srv/book/v1"
	"github.com/scylladb/scylla-cloud/service/book"
	"github.com/spf13/cobra"
)

func newBookCommand() *cobra.Command {
	cmd := serverCmd{
		Command: cobra.Command{
			Use:   "book",
			Short: "Book service",
		},
		SwaggerJSON:    pb.SwaggerJSON,
		ServiceHandler: pb.NewBookServiceServer(book.NewService()),
	}
	cmd.init()

	return &cmd.Command
}
