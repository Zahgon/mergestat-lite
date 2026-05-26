package enry

import (
	"go.riyazali.net/sqlite"
)

type EnryIsGenerated struct{}

func (f *EnryIsGenerated) Args() int           { _ = "STUB: not implemented"; return 0 }
func (f *EnryIsGenerated) Deterministic() bool { _ = "STUB: not implemented"; return false }
func (f *EnryIsGenerated) Apply(context *sqlite.Context, value ...sqlite.Value) {
	_ = "STUB: not implemented"
	return
}
