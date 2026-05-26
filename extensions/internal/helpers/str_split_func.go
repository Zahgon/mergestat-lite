package helpers

import (
	"go.riyazali.net/sqlite"
)

// StringSplit implements str_split scalar sql function.
// The function signature of the equivalent sql function is:
//
//	str_split(input, separator, index) string
type StringSplit struct{}

func (s *StringSplit) Args() int           { _ = "STUB: not implemented"; return 0 }
func (s *StringSplit) Deterministic() bool { _ = "STUB: not implemented"; return false }

func (s *StringSplit) Apply(context *sqlite.Context, value ...sqlite.Value) {
	_ = "STUB: not implemented"
	return
}
