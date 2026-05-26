package helpers

import (
	"go.riyazali.net/sqlite"
)

// TimeDiff implements a timediff pretty print function
// using github.com/mergestat/timediff
type TimeDiff struct{}

func (y *TimeDiff) Args() int           { _ = "STUB: not implemented"; return 0 }
func (y *TimeDiff) Deterministic() bool { _ = "STUB: not implemented"; return false }

func (y *TimeDiff) Apply(context *sqlite.Context, value ...sqlite.Value) {
	_ = "STUB: not implemented"
	return
}
