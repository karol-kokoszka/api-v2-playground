package cmd

import (
	mw "github.com/labstack/echo/v4/middleware"
	"github.com/spf13/cobra"
	"golang.org/x/time/rate"

	"github.com/scylladb/scylla-cloud/internal/restapi"
)

type apiServerCmd struct {
	serverCmd
	RateLimit int
	BookAddr  string
	ShelfAddr string
}

func newAPIServerCommand() *cobra.Command {
	cmd := apiServerCmd{
		serverCmd: serverCmd{
			Command: cobra.Command{
				Use:   "api",
				Short: "API server",
			},
			serverConfig: defaultServerConfig,
			//SwaggerJSON:  restapi.SwaggerJSON,
		},
		RateLimit: 100,
		BookAddr:  "http://book-svc:80",
		ShelfAddr: "http://shelf-svc:80",
	}
	cmd.HTTPAddr = ":80"
	cmd.init()

	return &cmd.Command
}

func (cmd *apiServerCmd) init() {
	fs := cmd.Flags()
	fs.IntVar(&cmd.RateLimit, "rate-limit", cmd.RateLimit, "max number of requests per IP address per second")
	fs.StringVar(&cmd.BookAddr, "book-addr", cmd.BookAddr, "twirp book endpoint")
	fs.StringVar(&cmd.ShelfAddr, "shelf-addr", cmd.ShelfAddr, "twirp shelf endpoint")

	cmd.PreRunE = func(_ *cobra.Command, _ []string) error {
		return cmd.createAndStartHTTPServer()
	}

	cmd.serverCmd.init(useMiddleware(
		mw.RateLimiter(mw.NewRateLimiterMemoryStore(rate.Limit(cmd.RateLimit))),
		mw.Secure(),
		mw.CORS(),
		// mw.CSRF(), it is disabled because Swagger UI needs to set header...
	))
}

func (cmd *apiServerCmd) createAndStartHTTPServer() (err error) {
	router := restapi.NewRouter()
	router.Run(":8080")
	return
}
