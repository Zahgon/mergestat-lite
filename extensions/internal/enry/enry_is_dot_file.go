package enry

import (
	"go.riyazali.net/sqlite"
)

type EnryIsDotFile struct{}

func (f *EnryIsDotFile) Args() int           { _ = "STUB: not implemented"; return 0 }
func (f *EnryIsDotFile) Deterministic() bool { _ = "STUB: not implemented"; return false }
func (f *EnryIsDotFile) Apply(context *sqlite.Context, value ...sqlite.Value) {
	_ = "STUB: not implemented"
	return
}
