package enry

import (
	"go.riyazali.net/sqlite"
)

type EnryDetectLanguage struct{}

func (f *EnryDetectLanguage) Args() int           { _ = "STUB: not implemented"; return 0 }
func (f *EnryDetectLanguage) Deterministic() bool { _ = "STUB: not implemented"; return false }
func (f *EnryDetectLanguage) Apply(context *sqlite.Context, value ...sqlite.Value) {
	_ = "STUB: not implemented"
	return
}
