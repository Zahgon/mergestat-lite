package native

import (
	"github.com/augmentable-dev/vtab"
	libgit2 "github.com/libgit2/git2go/v34"
	"github.com/mergestat/mergestat-lite/extensions/internal/git/utils"
	"go.riyazali.net/sqlite"
)

var blameCols = []vtab.Column{
	{Name: "line_no", Type: "TEXT", NotNull: false, Hidden: false, Filters: nil, OrderBy: vtab.NONE},
	{Name: "commit_hash", Type: "TEXT", NotNull: false, Hidden: false, Filters: nil, OrderBy: vtab.NONE},

	{Name: "repository", Type: "TEXT", NotNull: true, Hidden: true, Filters: []*vtab.ColumnFilter{{Op: sqlite.INDEX_CONSTRAINT_EQ, OmitCheck: true}}, OrderBy: vtab.NONE},
	{Name: "rev", Type: "TEXT", NotNull: true, Hidden: true, Filters: []*vtab.ColumnFilter{{Op: sqlite.INDEX_CONSTRAINT_EQ, OmitCheck: true}}, OrderBy: vtab.NONE},
	{Name: "file_path", Type: "TEXT", NotNull: true, Hidden: true, Filters: []*vtab.ColumnFilter{{Op: sqlite.INDEX_CONSTRAINT_EQ, OmitCheck: true}}, OrderBy: vtab.NONE},
}

// NewBlameModule returns the implementation of a table-valued-function for accessing git blame
func NewBlameModule(options *utils.ModuleOptions) sqlite.Module {
	_ = "STUB: not implemented"
	return *new(sqlite.Module)
}

func newBlameIter(options *utils.ModuleOptions, repoPath, rev, filePath string) (*blameIter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// if no rev is supplied, use HEAD

// TODO(patrickdevivo) figure out a better handling here

type blamedLine struct {
	lineNo int
	hunk   *libgit2.BlameHunk
}

type blameIter struct {
	repoPath string
	rev      string
	filePath string
	lines    []*blamedLine
	index    int
}

func (i *blameIter) Column(ctx vtab.Context, c int) error { _ = "STUB: not implemented"; return nil }

func (i *blameIter) Next() (vtab.Row, error) { _ = "STUB: not implemented"; return *new(vtab.Row), nil }
