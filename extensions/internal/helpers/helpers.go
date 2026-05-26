package helpers

import (
	"github.com/mergestat/mergestat-lite/extensions/options"
	"go.riyazali.net/sqlite"
)

// Register registers helpers as a SQLite extension
func Register(ext *sqlite.ExtensionApi, _ *options.Options) (_ sqlite.ErrorCode, err error) {
	_ = "STUB: not implemented"
	return *new(sqlite.ErrorCode), nil
}

// alias yaml_to_json => yml_to_json
