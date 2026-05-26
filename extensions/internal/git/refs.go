package git

import (
	"regexp"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/storer"
	"github.com/mergestat/mergestat-lite/extensions/internal/git/utils"
	"go.riyazali.net/sqlite"
)

var remoteName = regexp.MustCompile(`(?m)refs\/remotes\/([^\/]*)\/.+`)

// NewRefModule returns a new virtual table for listing git refs
func NewRefModule(opt *utils.ModuleOptions) sqlite.Module {
	_ = "STUB: not implemented"
	return *new(sqlite.Module)
}

type refModule struct {
	*utils.ModuleOptions
}

func (mod *refModule) Connect(_ *sqlite.Conn, _ []string, declare func(string) error) (sqlite.VirtualTable, error) {
	_ = "STUB: not implemented"
	return *new(sqlite.VirtualTable), nil
}

type gitRefTable struct {
	*utils.ModuleOptions
}

func (tab *gitRefTable) Disconnect() error { _ = "STUB: not implemented"; return nil }
func (tab *gitRefTable) Destroy() error    { _ = "STUB: not implemented"; return nil }
func (tab *gitRefTable) Open() (sqlite.VirtualCursor, error) {
	_ = "STUB: not implemented"
	return *new(sqlite.VirtualCursor), nil
}

func (tab *gitRefTable) BestIndex(input *sqlite.IndexInfoInput) (*sqlite.IndexInfoOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// if repository is provided, it must be usable

// we do not support unusable constraint at all

// validate passed in constraint to ensure there combination stays logical

type gitRefCursor struct {
	*utils.ModuleOptions

	repo *git.Repository

	ref  *plumbing.Reference
	refs storer.ReferenceIter
}

func (cur *gitRefCursor) Filter(_ int, s string, values ...sqlite.Value) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// values extracted from constraints

// open the git repository

func (cur *gitRefCursor) Column(c *sqlite.VirtualTableContext, col int) error {
	_ = "STUB: not implemented"
	return nil
}

func (cur *gitRefCursor) Next() (err error) { _ = "STUB: not implemented"; return nil }

func (cur *gitRefCursor) Eof() bool             { _ = "STUB: not implemented"; return false }
func (cur *gitRefCursor) Rowid() (int64, error) { _ = "STUB: not implemented"; return 0, nil }
func (cur *gitRefCursor) Close() error          { _ = "STUB: not implemented"; return nil }
