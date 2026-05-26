package github

import (
	"go.riyazali.net/sqlite"
)

// github_user_info
// This function takes in a user's login and returns user information as JSON
type userInfo struct {
	opts *Options
}

func (s *userInfo) Args() int           { _ = "STUB: not implemented"; return 0 }
func (s *userInfo) Deterministic() bool { _ = "STUB: not implemented"; return false }

func (s *userInfo) Apply(ctx *sqlite.Context, value ...sqlite.Value) {
	_ = "STUB: not implemented"
	return
}

func NewGitHubUserFunc(opts *Options) sqlite.Function {
	_ = "STUB: not implemented"
	return *new(sqlite.Function)
}
