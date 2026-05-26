// Package npm implements functions for iteracting with the npm registry
package npm

import (
	"context"
	"net/http"

	"github.com/mergestat/mergestat-lite/extensions/options"
	"github.com/rs/zerolog"
	"go.riyazali.net/sqlite"
)

const (
	BaseURL = "https://registry.npmjs.org"
)

type Client struct {
	httpClient *http.Client
	logger     *zerolog.Logger
}

// NewClient creates a new API client from an *http.Client. Pass nil to use http.DefaultClient
func NewClient(httpClient *http.Client, logger *zerolog.Logger) *Client {
	_ = "STUB: not implemented"
	return nil
}

// GetPackage makes an HTTP request to https://registry.npmjs.org/<<packageName>> and returns the JSON response
func (c *Client) GetPackage(ctx context.Context, packageName string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetPackageVersion makes an HTTP request to https://registry.npmjs.org/<<packageName>>/<<version>> and returns the JSON response
func (c *Client) GetPackageVersion(ctx context.Context, packageName, version string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Register registers npm API related functionality as a SQLite extension
func Register(ext *sqlite.ExtensionApi, opt *options.Options) (_ sqlite.ErrorCode, err error) {
	_ = "STUB: not implemented"
	return *new(sqlite.ErrorCode), nil
}
