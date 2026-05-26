package enry

import (
	"go.riyazali.net/sqlite"
)

type EnryIsTest struct{}

func (f *EnryIsTest) Args() int           { _ = "STUB: not implemented"; return 0 }
func (f *EnryIsTest) Deterministic() bool { _ = "STUB: not implemented"; return false }
func (f *EnryIsTest) Apply(context *sqlite.Context, value ...sqlite.Value) {
	_ = "STUB: not implemented"
	return
}
