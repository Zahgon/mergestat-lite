package enry

import (
	"go.riyazali.net/sqlite"
)

type EnryIsDocumentation struct{}

func (f *EnryIsDocumentation) Args() int           { _ = "STUB: not implemented"; return 0 }
func (f *EnryIsDocumentation) Deterministic() bool { _ = "STUB: not implemented"; return false }
func (f *EnryIsDocumentation) Apply(context *sqlite.Context, value ...sqlite.Value) {
	_ = "STUB: not implemented"
	return
}
