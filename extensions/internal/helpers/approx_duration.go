package helpers

import (
	"go.riyazali.net/sqlite"
)

// ApproxDuration pretty prints a duration given in days showing
// x years (y months)
// x years
// x months (y days)
// x months
// x days
type ApproxDuration struct{}

func (y *ApproxDuration) Args() int           { _ = "STUB: not implemented"; return 0 }
func (y *ApproxDuration) Deterministic() bool { _ = "STUB: not implemented"; return false }

func (y *ApproxDuration) Apply(context *sqlite.Context, value ...sqlite.Value) {
	_ = "STUB: not implemented"
	return
}
