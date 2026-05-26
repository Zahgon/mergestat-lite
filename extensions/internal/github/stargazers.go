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

type stargazer struct {
	Login           string
	Email           string
	Name            string
	Bio             string
	Company         string
	AvatarUrl       string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	TwitterUsername string
	WebsiteUrl      string
	Location        string
}

type stargazerEdge struct {
	StarredAt string
	Node      stargazer
}

type fetchStarsResults struct {
	RateLimit   *options.GitHubRateLimitResponse
	Edges       []*stargazerEdge
	HasNextPage bool
	EndCursor   *githubv4.String
}

func (i *iterStargazers) fetchStars(ctx context.Context, startCursor *githubv4.String) (*fetchStarsResults, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type iterStargazers struct {
	*Options
	owner     string
	name      string
	current   int
	results   *fetchStarsResults
	starOrder *githubv4.StarOrder
}

func (i *iterStargazers) logger() *zerolog.Logger { _ = "STUB: not implemented"; return nil }

func (i *iterStargazers) Column(ctx vtab.Context, c int) error {
	_ = "STUB: not implemented"
	return nil
}

func (i *iterStargazers) Next() (vtab.Row, error) {
	_ = "STUB: not implemented"
	return *new(vtab.Row), nil
}

var stargazersCols = []vtab.Column{
	{Name: "owner", Type: "TEXT", Hidden: true, Filters: []*vtab.ColumnFilter{{Op: sqlite.INDEX_CONSTRAINT_EQ, OmitCheck: true}}},
	{Name: "reponame", Type: "TEXT", Hidden: true, Filters: []*vtab.ColumnFilter{{Op: sqlite.INDEX_CONSTRAINT_EQ, OmitCheck: true}}},
	{Name: "login", Type: "TEXT"},
	{Name: "email", Type: "TEXT"},
	{Name: "name", Type: "TEXT"},
	{Name: "bio", Type: "TEXT"},
	{Name: "company", Type: "TEXT"},
	{Name: "avatar_url", Type: "TEXT"},
	{Name: "created_at", Type: "DATETIME"},
	{Name: "updated_at", Type: "DATETIME"},
	{Name: "twitter", Type: "TEXT"},
	{Name: "website", Type: "TEXT"},
	{Name: "location", Type: "TEXT"},
	{Name: "starred_at", Type: "DATETIME", OrderBy: vtab.ASC | vtab.DESC, Filters: []*vtab.ColumnFilter{
		{Op: sqlite.INDEX_CONSTRAINT_GT}, {Op: sqlite.INDEX_CONSTRAINT_GE},
		{Op: sqlite.INDEX_CONSTRAINT_LT}, {Op: sqlite.INDEX_CONSTRAINT_LE},
	}},
}

func NewStargazersModule(opts *Options) sqlite.Module {
	_ = "STUB: not implemented"
	return *new(sqlite.Module)
}

// for now we can only support single field order bys
