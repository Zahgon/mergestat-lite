package github

import (
	"context"

	"github.com/augmentable-dev/vtab"
	"github.com/mergestat/mergestat-lite/extensions/options"
	"github.com/rs/zerolog"
	"github.com/shurcooL/githubv4"
	"go.riyazali.net/sqlite"
)

type repositoryDefaultBranchForCommits struct {
	Target struct {
		Commits struct {
			History struct {
				Nodes    []*objectCommit
				PageInfo struct {
					EndCursor   githubv4.String
					HasNextPage bool
				}
			} `graphql:"history(first:100,after: $commitObjectCursor)"`
		} `graphql:"... on Commit"`
	}
}

type repositoryForCommits struct {
	Commits struct {
		History struct {
			Nodes    []*objectCommit
			PageInfo struct {
				EndCursor   githubv4.String
				HasNextPage bool
			}
		} `graphql:"history(first:100,after: $commitObjectCursor)"`
	} `graphql:"... on Commit"`
}

type objectCommit struct {
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

type fetchRepositoryCommitsResults struct {
	RateLimit      *options.GitHubRateLimitResponse
	defaultCommits *repositoryDefaultBranchForCommits
	Commits        *repositoryForCommits
	HasNextPage    bool
	EndCursor      *githubv4.String
}

func (i *iterRepositoryCommits) fetchRepositoryCommits(ctx context.Context, endCursor *githubv4.String) (*fetchRepositoryCommitsResults, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// if branch is passed instantiate it into variables and call the associated struct

type iterRepositoryCommits struct {
	*Options
	owner         string
	name          string
	branch        string
	currentCommit int
	results       *fetchRepositoryCommitsResults
}

func (i *iterRepositoryCommits) logger() *zerolog.Logger { _ = "STUB: not implemented"; return nil }

func (i *iterRepositoryCommits) Column(ctx vtab.Context, c int) error {
	_ = "STUB: not implemented"
	return nil

	// set current to whichever struct was used ie branch/ref is not passed vs it is
}

func (i *iterRepositoryCommits) Next() (vtab.Row, error) {
	_ = "STUB: not implemented"
	return *new(vtab.Row), nil
}

// set current to whichever struct was used ie branch/ref is not passed vs it is

var repoCommitCols = []vtab.Column{
	{Name: "owner", Type: "TEXT", NotNull: true, Hidden: true, Filters: []*vtab.ColumnFilter{{Op: sqlite.INDEX_CONSTRAINT_EQ, OmitCheck: true}}},
	{Name: "reponame", Type: "TEXT", NotNull: true, Hidden: true, Filters: []*vtab.ColumnFilter{{Op: sqlite.INDEX_CONSTRAINT_EQ, OmitCheck: true}}},

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

func NewRepoCommitsModule(opts *Options) sqlite.Module {
	_ = "STUB: not implemented"
	return *new(sqlite.Module)
}
