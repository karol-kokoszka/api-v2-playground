package cmd

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	mw "github.com/labstack/echo/v4/middleware"
	"github.com/scylladb/go-log"
)

func useMiddleware(middleware ...echo.MiddlewareFunc) func(e *echo.Echo) {
	return func(e *echo.Echo) {
		e.Use(middleware...)
	}
}

const dumpReqMsg = `request %s details:
---[ REQUEST ]---------------------------------------
%s
%s
`

const dumpReqErrMsg = `ERROR: request %s:
---[ REQUEST DUMP ERROR ]-----------------------------
%s
`

const dumpRespMsg = `---[ RESPONSE ]--------------------------------------
%s
%s
-----------------------------------------------------
`

const dumpRespErrMsg = `---[ RESPONSE DUMP ERROR ]-----------------------------
%s
-----------------------------------------------------
`

func bodyDumpMiddleware(log log.Logger) echo.MiddlewareFunc {
	return mw.BodyDumpWithConfig(mw.BodyDumpConfig{
		Skipper: func(c echo.Context) bool {
			return strings.HasPrefix(c.Request().URL.Path, "/ui/")
		},
		Handler: func(c echo.Context, reqBody, respBody []byte) {
			var b bytes.Buffer
			dumpRequest(&b, c.Request(), reqBody)
			dumpResponse(&b, c.Response(), respBody)
			log.Debug(c.Request().Context(), b.String())
		},
	})
}

func dumpRequest(b *bytes.Buffer, req *http.Request, body []byte) {
	r, err := httputil.DumpRequest(req, false)
	if err != nil {
		fmt.Fprintf(b, dumpReqErrMsg, req.URL.RequestURI(), err)
		return
	}

	fmt.Fprintf(b, dumpReqMsg, req.URL.RequestURI(), string(r), string(body))
}

func dumpResponse(b *bytes.Buffer, resp *echo.Response, body []byte) {
	if resp == nil {
		fmt.Fprintf(b, dumpRespErrMsg, "nil response")
		return
	}
	var h bytes.Buffer
	if err := resp.Header().Write(&h); err != nil {
		fmt.Fprintf(b, dumpRespErrMsg, err)
		return
	}

	fmt.Fprintf(b, dumpRespMsg, h.String(), string(body))
}

// requestLogger returns an echo middleware which logs incoming requests to the provided logger at the provided level.
func requestLogger(logger log.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			resp := c.Response()
			start := time.Now()

			if err := next(c); err != nil {
				c.Error(err)
			}

			logger.Info(req.Context(), "req",
				"from", req.RemoteAddr,
				"method", req.Method,
				"uri", req.RequestURI,
				"duration", time.Since(start),
				"status", resp.Status,
			)
			return nil
		}
	}
}
