package github

import (
	"go.riyazali.net/sqlite"
)

type repoInfo struct {
	opts *Options
}

func (r *repoInfo) Args() int           { _ = "STUB: not implemented"; return 0 }
func (r *repoInfo) Deterministic() bool { _ = "STUB: not implemented"; return false }
func (r *repoInfo) Apply(ctx *sqlite.Context, values ...sqlite.Value) {
	_ = "STUB: not implemented"
	return
}

// TODO(patrickdevivo) seems silly to unmarshal the graphql query into a struct
// and then immediately re-marshal it into json for output...should probably make this simpler
// by using the graphQL query directly and returning the []byte result directly

func NewRepoInfoFunc(opts *Options) sqlite.Function {
	_ = "STUB: not implemented"
	return *new(sqlite.Function)
}
