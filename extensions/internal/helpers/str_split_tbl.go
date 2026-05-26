package helpers

import (
	"bufio"

	"github.com/augmentable-dev/vtab"
	"go.riyazali.net/sqlite"
)

var strSplitCols = []vtab.Column{
	{Name: "line_no", Type: "INT", NotNull: false, Hidden: false, Filters: nil, OrderBy: vtab.NONE},
	{Name: "line", Type: "TEXT", NotNull: false, Hidden: false, Filters: nil, OrderBy: vtab.NONE},

	{Name: "contents", Type: "TEXT", NotNull: true, Hidden: true, Filters: []*vtab.ColumnFilter{{Op: sqlite.INDEX_CONSTRAINT_EQ, OmitCheck: true}}, OrderBy: vtab.NONE},
	{Name: "delimiter", Type: "TEXT", NotNull: true, Hidden: true, Filters: []*vtab.ColumnFilter{{Op: sqlite.INDEX_CONSTRAINT_EQ, OmitCheck: true}}, OrderBy: vtab.NONE},
}

// NewStrSplitModule returns the implementation of a table-valued-function for splitting string contents on a delimiter (default "\n")
func NewStrSplitModule() sqlite.Module { _ = "STUB: not implemented"; return *new(sqlite.Module) }

func newStrSplitIter(contents string, delimiter string) (*strSplitAllIter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// if a delimiter is provided, see here: https://stackoverflow.com/questions/33068644/how-a-scanner-can-be-implemented-with-a-custom-split/33069759

// Return nothing if at end of file and no data passed

// Find the index of the delimiter

// If at end of file with data return the data

// TODO make the buffer size settable

type strSplitAllIter struct {
	contents bufio.Scanner
	index    int
}

func (i *strSplitAllIter) Column(ctx vtab.Context, c int) error {
	_ = "STUB: not implemented"
	return nil
}

func (i *strSplitAllIter) Next() (vtab.Row, error) {
	_ = "STUB: not implemented"
	return *new(vtab.Row), nil
}
