package github

import (
	"context"

	"github.com/augmentable-dev/vtab"
	"github.com/mergestat/mergestat-lite/extensions/options"
	"github.com/rs/zerolog"
	"github.com/shurcooL/githubv4"
	"go.riyazali.net/sqlite"
)

type ref struct {
	Name   githubv4.String
	Prefix githubv4.String
	Target struct {
		Commit struct {
			Oid    githubv4.GitObjectID
			Author struct {
				Name  githubv4.String
				Email githubv4.String
			}
		} `graphql:"... on Commit"`
	}
}

type fetchBranchResults struct {
	RateLimit   *options.GitHubRateLimitResponse
	Edges       []*ref
	HasNextPage bool
	EndCursor   *githubv4.String
}

func (i *iterBranches) fetchBranches(ctx context.Context, startCursor *githubv4.String) (*fetchBranchResults, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type iterBranches struct {
	*Options
	owner   string
	name    string
	current int
	results *fetchBranchResults
}

func (i *iterBranches) logger() *zerolog.Logger { _ = "STUB: not implemented"; return nil }

func (i *iterBranches) Column(ctx vtab.Context, c int) error { _ = "STUB: not implemented"; return nil }

func (i *iterBranches) Next() (vtab.Row, error) {
	_ = "STUB: not implemented"
	return *new(vtab.Row), nil
}

var branchCols = []vtab.Column{
	{Name: "owner", Type: "TEXT", NotNull: true, Hidden: true, Filters: []*vtab.ColumnFilter{{Op: sqlite.INDEX_CONSTRAINT_EQ, OmitCheck: true}}},
	{Name: "reponame", Type: "TEXT", NotNull: true, Hidden: true, Filters: []*vtab.ColumnFilter{{Op: sqlite.INDEX_CONSTRAINT_EQ, OmitCheck: true}}},
	{Name: "name", Type: "TEXT"},
	{Name: "author_name", Type: "TEXT"},
	{Name: "author_email", Type: "TEXT"},
	{Name: "commit_hash", Type: "TEXT"},
}

func NewBranchModule(opts *Options) sqlite.Module {
	_ = "STUB: not implemented"
	return *new(sqlite.Module)
}
