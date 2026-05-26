package github

import (
	"github.com/mergestat/mergestat-lite/extensions/options"
	"go.riyazali.net/sqlite"
)

// Register registers GitHub related functionality as a SQLite extension
func Register(ext *sqlite.ExtensionApi, opt *options.Options) (_ sqlite.ErrorCode, err error) {
	_ = "STUB: not implemented"
	return *new(sqlite.ErrorCode), nil
}

// for now, just log to debug output current status of rate limit

// register GitHub tables

// register GitHub funcs
