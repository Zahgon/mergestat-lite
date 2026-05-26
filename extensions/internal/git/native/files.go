package native

import (
	"github.com/augmentable-dev/vtab"
	libgit2 "github.com/libgit2/git2go/v34"
	"github.com/mergestat/mergestat-lite/extensions/internal/git/utils"
	"go.riyazali.net/sqlite"
)

var filesCols = []vtab.Column{
	{Name: "path", Type: "TEXT", NotNull: false, Hidden: false, Filters: nil, OrderBy: vtab.NONE},
	{Name: "executable", Type: "INT", NotNull: false, Hidden: false, Filters: nil, OrderBy: vtab.NONE},
	{Name: "contents", Type: "BLOB", NotNull: false, Hidden: false, Filters: nil, OrderBy: vtab.NONE},

	{Name: "repository", Type: "TEXT", NotNull: true, Hidden: true, Filters: []*vtab.ColumnFilter{{Op: sqlite.INDEX_CONSTRAINT_EQ, OmitCheck: true}}, OrderBy: vtab.NONE},
	{Name: "rev", Type: "TEXT", NotNull: true, Hidden: true, Filters: []*vtab.ColumnFilter{{Op: sqlite.INDEX_CONSTRAINT_EQ, OmitCheck: true}}, OrderBy: vtab.NONE},
}

// NewFilesModule returns the implementation of a table-valued-function for accessing the content of files in git
func NewFilesModule(options *utils.ModuleOptions) sqlite.Module {
	_ = "STUB: not implemented"
	return *new(sqlite.Module)
}

func newFilesIter(options *utils.ModuleOptions, repoPath, rev string) (*filesIter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// if no rev is supplied, use HEAD

type file struct {
	id         *libgit2.Oid
	path       string
	executable bool
}

type filesIter struct {
	repoPath string
	rev      string
	files    []*file
	index    int
	repo     *libgit2.Repository
}

func (i *filesIter) Column(ctx vtab.Context, c int) error { _ = "STUB: not implemented"; return nil }

func (i *filesIter) Next() (vtab.Row, error) { _ = "STUB: not implemented"; return *new(vtab.Row), nil }
