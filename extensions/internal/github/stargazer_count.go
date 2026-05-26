package github

import (
	"go.riyazali.net/sqlite"
)

type starCount struct {
	opts *Options
}

func (s *starCount) Args() int           { _ = "STUB: not implemented"; return 0 }
func (s *starCount) Deterministic() bool { _ = "STUB: not implemented"; return false }
func (s *starCount) Apply(ctx *sqlite.Context, values ...sqlite.Value) {
	_ = "STUB: not implemented"
	return
}

func NewStarredReposFunc(opts *Options) sqlite.Function {
	_ = "STUB: not implemented"
	return *new(sqlite.Function)
}
