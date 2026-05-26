package enry

import (
	"go.riyazali.net/sqlite"
)

type EnryIsConfiguration struct{}

func (f *EnryIsConfiguration) Args() int           { _ = "STUB: not implemented"; return 0 }
func (f *EnryIsConfiguration) Deterministic() bool { _ = "STUB: not implemented"; return false }
func (f *EnryIsConfiguration) Apply(context *sqlite.Context, value ...sqlite.Value) {
	_ = "STUB: not implemented"
	return
}
