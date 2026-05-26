package enry

import (
	"github.com/mergestat/mergestat-lite/extensions/options"
	"go.riyazali.net/sqlite"
)

// Register registers enry related functionality as a SQLite extension
func Register(ext *sqlite.ExtensionApi, _ *options.Options) (_ sqlite.ErrorCode, err error) {
	_ = "STUB: not implemented"
	return *new(sqlite.ErrorCode), nil
}
