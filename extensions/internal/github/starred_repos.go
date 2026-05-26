package github

import (
	"context"
	"time"

	"github.com/augmentable-dev/vtab"
	"github.com/mergestat/mergestat-lite/extensions/options"
	"github.com/rs/zerolog"
	"github.com/shurcooL/githubv4"
	"go.riyazali.net/sqlite"
)

type fetchStarredReposResults struct {
	RateLimit   *options.GitHubRateLimitResponse
	Edges       []*starredRepoEdge
	HasNextPage bool
	EndCursor   *githubv4.String
}

type starredRepoEdge struct {
	StarredAt string
	Node      *starredRepoNode
}

type starredRepoNode struct {
	Name           string
	Url            string
	Description    string
	CreatedAt      time.Time
	PushedAt       time.Time
	UpdatedAt      time.Time
	StargazerCount int
	NameWithOwner  string
}

func (i *iterStarredRepos) fetchStarredRepos(ctx context.Context, startCursor *githubv4.String) (*fetchStarredReposResults, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type iterStarredRepos struct {
	*Options
	login     string
	current   int
	results   *fetchStarredReposResults
	starOrder *githubv4.StarOrder
}

func (i *iterStarredRepos) logger() *zerolog.Logger { _ = "STUB: not implemented"; return nil }

func (i *iterStarredRepos) Column(ctx vtab.Context, c int) error {
	_ = "STUB: not implemented"
	return nil
}

func (i *iterStarredRepos) Next() (vtab.Row, error) {
	_ = "STUB: not implemented"
	return *new(vtab.Row), nil
}

var starredReposCols = []vtab.Column{
	{Name: "login", Type: "TEXT", NotNull: false, Hidden: true, Filters: []*vtab.ColumnFilter{{Op: sqlite.INDEX_CONSTRAINT_EQ, OmitCheck: true}}},
	{Name: "name", Type: "TEXT"},
	{Name: "url", Type: "TEXT"},
	{Name: "description", Type: "TEXT"},
	{Name: "created_at", Type: "DATETIME"},
	{Name: "pushed_at", Type: "DATETIME"},
	{Name: "updated_at", Type: "DATETIME"},
	{Name: "stargazer_count", Type: "INT"},
	{Name: "name_with_owner", Type: "TEXT"},
	{Name: "starred_at", Type: "DATETIME", OrderBy: vtab.ASC | vtab.DESC},
}

func NewStarredReposModule(opts *Options) sqlite.Module {
	_ = "STUB: not implemented"
	return *new(sqlite.Module)
}

// for now we can only support single field order bys
