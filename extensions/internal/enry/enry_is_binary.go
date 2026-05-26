package enry

import (
	"go.riyazali.net/sqlite"
)

type EnryIsBinary struct{}

func (f *EnryIsBinary) Args() int           { _ = "STUB: not implemented"; return 0 }
func (f *EnryIsBinary) Deterministic() bool { _ = "STUB: not implemented"; return false }
func (f *EnryIsBinary) Apply(context *sqlite.Context, value ...sqlite.Value) {
	_ = "STUB: not implemented"
	return
}
