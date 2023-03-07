package cmd

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	mw "github.com/labstack/echo/v4/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/scylladb/go-log"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/scylladb/scylla-cloud/util/flagtype"
	"github.com/scylladb/scylla-cloud/util/httppprof"
)

type serverConfig struct {
	HTTPAddr   string
	HTTPSAddr  string
	CertFile   string
	KeyFile    string
	PprofAddr  string
	PromAddr   string
	BodyLimit  string
	ReqTimeout time.Duration
	Debug      bool
	LogLevel   string
}

var defaultServerConfig = serverConfig{
	HTTPAddr:   ":80",
	PprofAddr:  ":5112",
	PromAddr:   ":5090",
	BodyLimit:  "2M",
	ReqTimeout: 30 * time.Second,
	LogLevel:   "INFO",
}

func defaultReadiness(c echo.Context) error {
	return c.String(http.StatusOK, "ready")
}

type serverCmd struct {
	cobra.Command
	serverConfig

	SwaggerJSON    []byte
	ServiceHandler http.Handler
	ReadinessFunc  echo.HandlerFunc

	logger log.Logger
}

type echoOpt func(*echo.Echo)

func (cmd *serverCmd) init(opts ...echoOpt) {
	zero := serverConfig{}
	if cmd.serverConfig == zero {
		cmd.serverConfig = defaultServerConfig
	}

	if cmd.ReadinessFunc == nil {
		cmd.ReadinessFunc = defaultReadiness
	}

	fs := cmd.Flags()
	fs.StringVar(&cmd.HTTPAddr, "http", cmd.HTTPAddr, "http address")
	fs.StringVar(&cmd.HTTPSAddr, "https", cmd.HTTPSAddr, "https address")
	fs.StringVar(&cmd.CertFile, "cert-file", cmd.CertFile, "tls certificate file")
	fs.StringVar(&cmd.KeyFile, "key-file", cmd.KeyFile, "tls key file")
	fs.StringVar(&cmd.PprofAddr, "pprof", cmd.PprofAddr, "http address of pprof debug endpoints")
	fs.StringVar(&cmd.PromAddr, "prom", cmd.PromAddr, "http address of prometheus metrics endpoint")
	fs.StringVar(&cmd.BodyLimit, "body-limit", cmd.BodyLimit, "max body size")
	fs.Var(flagtype.WrapDuration(&cmd.ReqTimeout), "req-timeout", "max request processing time")
	fs.BoolVar(&cmd.Debug, "debug", cmd.Debug, "enable debug")
	fs.StringVar(&cmd.LogLevel, "log", cmd.LogLevel, "log level")

	cmd.addEnvToUsage()

	cmd.RunE = func(_ *cobra.Command, _ []string) error {
		if err := cmd.initLogger(); err != nil {
			return err
		}
		return cmd.runEcho(cmd.newEcho(opts...))
	}
}

func (cmd *serverCmd) newEcho(opts ...echoOpt) *echo.Echo {
	p := prometheus.NewPrometheus("scylla_cloud_"+cmd.Name(), nil)
	p.RequestCounterURLLabelMappingFunc = func(c echo.Context) string {
		return c.Request().URL.EscapedPath()
	}

	l := cmd.logger.Named("http")
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.Debug = cmd.Debug
	e.Logger = newEchoLogger(l)

	// WARN: middlewares are applied from the end, additional middlewares are applied first
	for _, o := range opts {
		o(e)
	}
	e.Use(
		p.HandlerFunc,

		mw.Gzip(),
		mw.Decompress(),

		mw.BodyLimit(cmd.BodyLimit),
		mw.TimeoutWithConfig(mw.TimeoutConfig{
			Timeout: cmd.ReqTimeout,
		}),
		mw.Recover(),
	)
	if cmd.Debug {
		e.Use(bodyDumpMiddleware(l))
	} else {
		e.Use(requestLogger(l))
	}

	if cmd.HTTPSAddr != "" {
		e.TLSServer.TLSConfig = tlsConfig()
	}
	if len(cmd.SwaggerJSON) != 0 {
		e.GET("/ui/swagger.json", func(c echo.Context) error {
			return c.Blob(200, "application/json", cmd.SwaggerJSON)
		})
		//e.StaticFS("/ui", swagger.UI())
		e.Any("/", func(c echo.Context) error {
			return c.Redirect(http.StatusPermanentRedirect, "/ui")
		})
	}

	e.GET("/health/live", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	e.GET("/health/ready", cmd.ReadinessFunc)

	e.Any("/*", echo.WrapHandler(cmd.ServiceHandler))

	return e
}

func tlsConfig() *tls.Config {
	return &tls.Config{
		MinVersion: tls.VersionTLS13,
	}
}

func (cmd *serverCmd) runEcho(e *echo.Echo) error {
	ctx := context.Background()
	errCh := make(chan error)

	if cmd.HTTPSAddr != "" {
		cmd.logger.Info(ctx, "starting https server", "addr", cmd.HTTPSAddr)
		go func() {
			if cmd.CertFile == "" || cmd.KeyFile == "" {
				errCh <- fmt.Errorf("https: requires cert-file and key-file")
				return
			}

			if err := e.StartTLS(cmd.HTTPSAddr, cmd.CertFile, cmd.KeyFile); err != nil {
				errCh <- fmt.Errorf("https: %w", err)
			}
		}()
	}

	if cmd.HTTPAddr != "" {
		cmd.logger.Info(ctx, "starting http server", "addr", cmd.HTTPAddr)
		go func() {
			if err := e.Start(cmd.HTTPAddr); err != nil {
				errCh <- fmt.Errorf("http: %w", err)
			}
		}()
	}

	if cmd.PprofAddr != "" {
		cmd.logger.Info(ctx, "starting pprof server", "addr", cmd.PprofAddr)
		go func() {
			if err := http.ListenAndServe(cmd.PprofAddr, httppprof.Handler()); err != nil {
				errCh <- fmt.Errorf("pprof: %w", err)
			}
		}()
	}

	if cmd.PromAddr != "" {
		cmd.logger.Info(ctx, "starting prometheus server", "addr", cmd.PromAddr)
		go func() {
			if err := http.ListenAndServe(cmd.PromAddr, promhttp.Handler()); err != nil {
				errCh <- fmt.Errorf("prometheus: %w", err)
			}
		}()
	}

	err := <-errCh
	return err
}

func (cmd *serverCmd) initLogger() error {
	level, err := log.ParseLevel(cmd.LogLevel)
	if err != nil {
		return err
	}
	if cmd.Debug {
		level, _ = log.ParseLevel("DEBUG") // nolint: errcheck
	}

	cmd.logger = log.NewDevelopmentWithLevel(level.Level())

	return nil
}

func (cmd *serverCmd) addEnvToUsage() {
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		envVarSuffix := strings.ToUpper(strings.ReplaceAll(f.Name, "-", "_"))
		f.Usage = fmt.Sprintf("%s [ENV: %s_%s]", f.Usage, envPrefix, envVarSuffix)
	})
}
