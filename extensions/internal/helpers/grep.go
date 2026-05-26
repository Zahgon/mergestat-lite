package helpers

import (
	"regexp"

	"github.com/augmentable-dev/vtab"
	"go.riyazali.net/sqlite"
)

var grepCols = []vtab.Column{
	{Name: "line_no", Type: "INT", OrderBy: vtab.NONE},
	{Name: "line", Type: "TEXT", OrderBy: vtab.NONE},

	{Name: "contents", Type: "TEXT", Hidden: true, Filters: []*vtab.ColumnFilter{{Op: sqlite.INDEX_CONSTRAINT_EQ, OmitCheck: true}}, OrderBy: vtab.NONE},
	{Name: "search", Type: "TEXT", Hidden: true, Filters: []*vtab.ColumnFilter{{Op: sqlite.INDEX_CONSTRAINT_EQ, OmitCheck: true}}, OrderBy: vtab.NONE},
	{Name: "preceeding", Type: "INT", Hidden: true, Filters: []*vtab.ColumnFilter{{Op: sqlite.INDEX_CONSTRAINT_EQ, OmitCheck: true}}, OrderBy: vtab.NONE},
	{Name: "proceeding", Type: "INT", Hidden: true, Filters: []*vtab.ColumnFilter{{Op: sqlite.INDEX_CONSTRAINT_EQ, OmitCheck: true}}, OrderBy: vtab.NONE},
}

// NewStatsModule returns the implementation of a table-valued-function for grep
func NewGrepModule() sqlite.Module { _ = "STUB: not implemented"; return *new(sqlite.Module) }

// TODO(patrickdevivo) not entirely sure if we should fail/error on this or let it be

func newGrepIter(contents, search string, preceeding, proceeding int) (*grepIter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type grepIter struct {
	contents    string
	search      *regexp.Regexp
	preceeding  int
	proceeding  int
	splitString []string
	index       int
}

func (i *grepIter) Column(ctx vtab.Context, c int) error { _ = "STUB: not implemented"; return nil }

func (i *grepIter) Next() (vtab.Row, error) { _ = "STUB: not implemented"; return *new(vtab.Row), nil }
