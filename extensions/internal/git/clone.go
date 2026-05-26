package git

import (
	"github.com/mergestat/mergestat-lite/extensions/internal/git/utils"
	"go.riyazali.net/sqlite"
)

// CloneFn is essentially a no-op that's useful for cloning remote repos
// by opening them (and calling the Locator)
type CloneFn struct {
	Options *utils.ModuleOptions
}

// NewCloneFn returns a new CloneFn implementation
func NewCloneFn(opt *utils.ModuleOptions) *CloneFn { _ = "STUB: not implemented"; return nil }

func (*CloneFn) Deterministic() bool { _ = "STUB: not implemented"; return false }
func (*CloneFn) Args() int           { _ = "STUB: not implemented"; return 0 }
func (fn *CloneFn) Apply(c *sqlite.Context, values ...sqlite.Value) {
	_ = "STUB: not implemented"
	return
}
