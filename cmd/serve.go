package cmd

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

var (
	servicePort int
)

func init() {
	serveCmd.Flags().IntVarP(&servicePort, "port", "p", 8000, "port to listen on")
}

// ServiceQueryRequest is the JSON body from a query HTTP request
type ServiceQueryRequest struct {
	Query string `json:"query"`
}

type queryServiceHandler struct {
	DB *sql.DB
}

func newQueryServiceHandler() (*queryServiceHandler, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *queryServiceHandler) Close() error { _ = "STUB: not implemented"; return nil }

// handleErr is a helper for writing errors to the http response
func (h *queryServiceHandler) handleErr(w http.ResponseWriter, statusCode int, err error) {
	_ = "STUB: not implemented"
	return
}

func (h *queryServiceHandler) httpHandler(w http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	return
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run an HTTP API server for receiving queries to execute",
	Long:  `Use this command to start a query API server`,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		// TODO(patrickdevivo) we might want to figure out a better logger set-up here.
		// For instance, to separate HTTP request logs from SQL execution logs.
		// Right now, they are all mixed together and sent to the global logger.
		var srv *queryServiceHandler
		var err error
		if srv, err = newQueryServiceHandler(); err != nil {
			handleExitError(err)
		}
		defer func() {
			if err := srv.Close(); err != nil {
				handleExitError(err)
			}
		}()

		http.HandleFunc("/", srv.httpHandler)
		http.HandleFunc("/query", srv.httpHandler)

		logger.Info().Msgf("starting HTTP API server on port %d", servicePort)
		if err := http.ListenAndServe(fmt.Sprintf(":%d", servicePort), nil); err != nil {
			handleExitError(err)
		}
	},
}
