package github

import (
	"go.riyazali.net/sqlite"
)

type repoFileContent struct {
	opts *Options
}

func (f *repoFileContent) Args() int           { _ = "STUB: not implemented"; return 0 }
func (f *repoFileContent) Deterministic() bool { _ = "STUB: not implemented"; return false }
func (f *repoFileContent) Apply(ctx *sqlite.Context, values ...sqlite.Value) {
	_ = "STUB: not implemented"
	return
}

func NewRepoFileContentFunc(opts *Options) sqlite.Function {
	_ = "STUB: not implemented"
	return *new(sqlite.Function)
}
