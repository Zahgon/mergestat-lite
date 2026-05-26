package git

import (
	"github.com/go-git/go-git/v5/plumbing"
)

// returns true if error is an end-of-file error
func eof(err error) bool { _ = "STUB: not implemented"; return false }

func enc(buf []byte) string          { _ = "STUB: not implemented"; return "" }
func dec(str string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func isRemoteBranch(ref plumbing.ReferenceName) bool { _ = "STUB: not implemented"; return false }
