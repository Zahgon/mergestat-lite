package native

import (
	"github.com/augmentable-dev/vtab"
	"github.com/mergestat/mergestat-lite/extensions/internal/git/utils"
	"go.riyazali.net/sqlite"
)

var statsCols = []vtab.Column{
	{Name: "file_path", Type: "TEXT", NotNull: false, Hidden: false, Filters: nil, OrderBy: vtab.NONE},
	{Name: "additions", Type: "INT", NotNull: false, Hidden: false, Filters: nil, OrderBy: vtab.NONE},
	{Name: "deletions", Type: "INT", NotNull: false, Hidden: false, Filters: nil, OrderBy: vtab.NONE},

	{Name: "old_file_mode", Type: "TEXT", NotNull: true, Hidden: false, Filters: nil, OrderBy: vtab.NONE},
	{Name: "new_file_mode", Type: "TEXT", NotNull: true, Hidden: false, Filters: nil, OrderBy: vtab.NONE},

	{Name: "repository", Type: "TEXT", NotNull: true, Hidden: true, Filters: []*vtab.ColumnFilter{{Op: sqlite.INDEX_CONSTRAINT_EQ, OmitCheck: true}}, OrderBy: vtab.NONE},
	{Name: "rev", Type: "TEXT", NotNull: true, Hidden: true, Filters: []*vtab.ColumnFilter{{Op: sqlite.INDEX_CONSTRAINT_EQ, OmitCheck: true}}, OrderBy: vtab.NONE},
	{Name: "to_rev", Type: "TEXT", NotNull: true, Hidden: true, Filters: []*vtab.ColumnFilter{{Op: sqlite.INDEX_CONSTRAINT_EQ, OmitCheck: true}}, OrderBy: vtab.NONE},
}

type GitFileModeObjectType string

const (
	GitFileModeObjectTypeUnknown     GitFileModeObjectType = "unknown"
	GitFileModeObjectTypeNone        GitFileModeObjectType = "none"
	GitFileModeObjectTypeRegularFile GitFileModeObjectType = "regular_file"
	GitFileModeOjectTypeSymbolicLink GitFileModeObjectType = "symbolic_link"
	GitFileModeOjectTypeGitLink      GitFileModeObjectType = "git_link"
)

// gitFileModeObjectTypeFromUint16 takes a git stats file mode and returns the GitFileModeObjectType.
// See here for more info on the modes: https://unix.stackexchange.com/questions/450480/file-permission-with-six-bytes-in-git-what-does-it-mean
func gitFileModeObjectTypeFromUint16(mode uint16) GitFileModeObjectType {
	_ = "STUB: not implemented"
	return *new(GitFileModeObjectType)
}

// NewStatsModule returns the implementation of a table-valued-function for git stats
func NewStatsModule(options *utils.ModuleOptions) sqlite.Module {
	_ = "STUB: not implemented"
	return *new(sqlite.Module)
}

func newStatsIter(options *utils.ModuleOptions, repoPath, rev, toRev string) (*statsIter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// if no rev is supplied, use HEAD

// TODO what should we do here?

type stat struct {
	filePath    string
	additions   int
	deletions   int
	oldFileMode GitFileModeObjectType
	newFileMode GitFileModeObjectType
}

type statsIter struct {
	repoPath string
	stats    []*stat
	index    int
}

func (i *statsIter) Column(ctx vtab.Context, c int) error { _ = "STUB: not implemented"; return nil }

func (i *statsIter) Next() (vtab.Row, error) { _ = "STUB: not implemented"; return *new(vtab.Row), nil }
