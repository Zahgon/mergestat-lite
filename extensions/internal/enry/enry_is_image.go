package enry

import (
	"go.riyazali.net/sqlite"
)

type EnryIsImage struct{}

func (f *EnryIsImage) Args() int           { _ = "STUB: not implemented"; return 0 }
func (f *EnryIsImage) Deterministic() bool { _ = "STUB: not implemented"; return false }
func (f *EnryIsImage) Apply(context *sqlite.Context, value ...sqlite.Value) {
	_ = "STUB: not implemented"
	return
}
