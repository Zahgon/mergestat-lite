package enry

import (
	"go.riyazali.net/sqlite"
)

type EnryIsVendor struct{}

func (f *EnryIsVendor) Args() int           { _ = "STUB: not implemented"; return 0 }
func (f *EnryIsVendor) Deterministic() bool { _ = "STUB: not implemented"; return false }
func (f *EnryIsVendor) Apply(context *sqlite.Context, value ...sqlite.Value) {
	_ = "STUB: not implemented"
	return
}
