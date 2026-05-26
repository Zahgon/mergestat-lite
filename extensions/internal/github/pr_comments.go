package github

import (
	"context"

	"github.com/augmentable-dev/vtab"
	"github.com/mergestat/mergestat-lite/extensions/options"
	"github.com/rs/zerolog"
	"github.com/shurcooL/githubv4"
	"go.riyazali.net/sqlite"
)

type pullRequestForComments struct {
	Id       githubv4.String
	Number   int
	Comments struct {
		Nodes    []*prComment
		PageInfo struct {
			EndCursor   githubv4.String
			HasNextPage bool
		}
	} `graphql:"comments(first: $perPage, after: $commentcursor,orderBy: $orderBy)"`
}

type prComment struct {
	Body   string
	Author struct {
		Login string
		Url   string
	}
	CreatedAt  githubv4.DateTime
	DatabaseId int
	Id         githubv4.GitObjectID
	UpdatedAt  githubv4.DateTime
	Url        githubv4.URI
}

type fetchPRCommentsResults struct {
	RateLimit   *options.GitHubRateLimitResponse
	Comments    *pullRequestForComments
	OrderBy     *githubv4.IssueCommentOrder
	HasNextPage bool
	EndCursor   *githubv4.String
}

func (i *iterPRComments) fetchPRComments(ctx context.Context, endCursor *githubv4.String) (*fetchPRCommentsResults, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type iterPRComments struct {
	*Options
	owner          string
	name           string
	prNumber       int
	currentComment int
	orderBy        *githubv4.IssueCommentOrder
	results        *fetchPRCommentsResults
}

func (i *iterPRComments) logger() *zerolog.Logger { _ = "STUB: not implemented"; return nil }

func (i *iterPRComments) Column(ctx vtab.Context, c int) error {
	_ = "STUB: not implemented"
	return nil
}

func (i *iterPRComments) Next() (vtab.Row, error) {
	_ = "STUB: not implemented"
	return *new(vtab.Row), nil
}

var prCommentCols = []vtab.Column{
	{Name: "owner", Type: "TEXT", NotNull: true, Hidden: true, Filters: []*vtab.ColumnFilter{{Op: sqlite.INDEX_CONSTRAINT_EQ, OmitCheck: true}}},
	{Name: "reponame", Type: "TEXT", NotNull: true, Hidden: true, Filters: []*vtab.ColumnFilter{{Op: sqlite.INDEX_CONSTRAINT_EQ, OmitCheck: true}}},
	{Name: "pr_number", Type: "INT", NotNull: true, Hidden: true, Filters: []*vtab.ColumnFilter{{Op: sqlite.INDEX_CONSTRAINT_EQ, OmitCheck: true}}},
	{Name: "author_login", Type: "TEXT"},
	{Name: "author_url", Type: "TEXT"},
	{Name: "body", Type: "TEXT"},
	{Name: "created_at", Type: "TEXT"},
	{Name: "database_id", Type: "INT"},
	{Name: "id", Type: "TEXT"},
	{Name: "updated_at", Type: "TEXT", OrderBy: vtab.ASC | vtab.DESC},
	{Name: "url", Type: "TEXT"},
	{Name: "pr_id", Type: "TEXT"},
}

func NewPRCommentsModule(opts *Options) sqlite.Module {
	_ = "STUB: not implemented"
	return *new(sqlite.Module)
}

// if true, user supplied 3 args, 1st is org name, 2nd is repo name, 3rd is pr number
