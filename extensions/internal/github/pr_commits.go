package github

import (
	"context"

	"github.com/augmentable-dev/vtab"
	"github.com/mergestat/mergestat-lite/extensions/options"
	"github.com/rs/zerolog"
	"github.com/shurcooL/githubv4"
	"go.riyazali.net/sqlite"
)

type pullRequestForCommits struct {
	Id      githubv4.String
	Number  int
	Commits struct {
		Nodes []struct {
			Commit *prCommit
		}
		PageInfo struct {
			EndCursor   githubv4.String
			HasNextPage bool
		}
	} `graphql:"commits(first: $perPage, after: $commitcursor)"`
}

type prCommit struct {
	Additions int
	Author    struct {
		Email string
		Name  string
		Date  githubv4.DateTime
	}
	ChangedFiles int
	Committer    struct {
		Email string
		Name  string
		Date  githubv4.DateTime
	}
	Deletions int
	Oid       githubv4.GitObjectID
	Message   string
	Url       githubv4.URI
}

type fetchPRCommitsResults struct {
	RateLimit   *options.GitHubRateLimitResponse
	PR          *pullRequestForCommits
	HasNextPage bool
	EndCursor   *githubv4.String
}

func (i *iterPRCommits) fetchPRCommits(ctx context.Context, endCursor *githubv4.String) (*fetchPRCommitsResults, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type iterPRCommits struct {
	*Options
	owner         string
	name          string
	prNumber      int
	currentCommit int
	results       *fetchPRCommitsResults
}

func (i *iterPRCommits) logger() *zerolog.Logger { _ = "STUB: not implemented"; return nil }

func (i *iterPRCommits) Column(ctx vtab.Context, c int) error {
	_ = "STUB: not implemented"
	return nil
}

func (i *iterPRCommits) Next() (vtab.Row, error) {
	_ = "STUB: not implemented"
	return *new(vtab.Row), nil
}

var prCommitsCols = []vtab.Column{
	{Name: "owner", Type: "TEXT", NotNull: true, Hidden: true, Filters: []*vtab.ColumnFilter{{Op: sqlite.INDEX_CONSTRAINT_EQ, OmitCheck: true}}},
	{Name: "reponame", Type: "TEXT", NotNull: true, Hidden: true, Filters: []*vtab.ColumnFilter{{Op: sqlite.INDEX_CONSTRAINT_EQ, OmitCheck: true}}},
	{Name: "pr_number", Type: "INT", NotNull: true, Hidden: true, Filters: []*vtab.ColumnFilter{{Op: sqlite.INDEX_CONSTRAINT_EQ, OmitCheck: true}}},

	{Name: "hash", Type: "TEXT"},
	{Name: "message", Type: "TEXT"},
	{Name: "author_name", Type: "TEXT"},
	{Name: "author_email", Type: "TEXT"},
	{Name: "author_when", Type: "DATETIME"},
	{Name: "committer_name", Type: "TEXT"},
	{Name: "committer_email", Type: "TEXT"},
	{Name: "committer_when", Type: "DATETIME"},
	{Name: "additions", Type: "INT"},
	{Name: "deletions", Type: "INT"},
	{Name: "changed_files", Type: "INT"},
	{Name: "url", Type: "TEXT"},
}

func NewPRCommitsModule(opts *Options) sqlite.Module {
	_ = "STUB: not implemented"
	return *new(sqlite.Module)
}

// if true, user supplied 3 args, 1st is org name, 2nd is repo name, 3rd is pr number
