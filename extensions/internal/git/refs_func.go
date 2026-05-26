package git

import (
	"go.riyazali.net/sqlite"
)

// CommitFromTagFn implements the COMMIT_FROM_TAG(...) sql function
type CommitFromTagFn struct{}

func (*CommitFromTagFn) Deterministic() bool { _ = "STUB: not implemented"; return false }
func (*CommitFromTagFn) Args() int           { _ = "STUB: not implemented"; return 0 }
func (*CommitFromTagFn) Apply(c *sqlite.Context, values ...sqlite.Value) {
	_ = "STUB: not implemented"
	return
}
