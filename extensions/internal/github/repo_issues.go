package github

import (
	"context"

	"github.com/augmentable-dev/vtab"
	"github.com/mergestat/mergestat-lite/extensions/options"
	"github.com/rs/zerolog"
	"github.com/shurcooL/githubv4"
	"go.riyazali.net/sqlite"
)

type issue struct {
	Author struct {
		Login string
	}
	Body     string
	Closed   bool
	ClosedAt githubv4.DateTime
	Comments struct {
		TotalCount int
	}
	CreatedAt       githubv4.DateTime
	CreatedViaEmail bool
	DatabaseId      int
	Editor          struct {
		Login string
	}
	IncludesCreatedEdit bool
	IsReadByViewer      bool
	Labels              struct {
		TotalCount int
		Nodes      []struct {
			Name string
		}
	} `graphql:"labels(first: 15)"`
	LastEditedAt githubv4.DateTime
	Locked       bool
	Milestone    struct {
		Number int
	}
	Number       int
	Participants struct {
		TotalCount int
	}
	PublishedAt githubv4.DateTime
	Reactions   struct {
		TotalCount int
	}
	State     githubv4.IssueState
	Title     string
	UpdatedAt githubv4.DateTime
	Url       githubv4.URI
}

type fetchIssuesResults struct {
	RateLimit   *options.GitHubRateLimitResponse
	Edges       []*issueEdge
	HasNextPage bool
	EndCursor   *githubv4.String
}

type issueEdge struct {
	Cursor string
	Node   issue
}

func (i *iterIssues) fetchIssues(ctx context.Context, startCursor *githubv4.String) (*fetchIssuesResults, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type iterIssues struct {
	*Options
	owner      string
	name       string
	current    int
	results    *fetchIssuesResults
	issueOrder *githubv4.IssueOrder
}

func (i *iterIssues) logger() *zerolog.Logger { _ = "STUB: not implemented"; return nil }

func (i *iterIssues) Column(ctx vtab.Context, c int) error { _ = "STUB: not implemented"; return nil }

func (i *iterIssues) Next() (vtab.Row, error) {
	_ = "STUB: not implemented"
	return *new(vtab.Row), nil
}

var issuesCols = []vtab.Column{
	{Name: "owner", Type: "TEXT", NotNull: true, Hidden: true, Filters: []*vtab.ColumnFilter{{Op: sqlite.INDEX_CONSTRAINT_EQ, OmitCheck: true}}},
	{Name: "reponame", Type: "TEXT", NotNull: true, Hidden: true, Filters: []*vtab.ColumnFilter{{Op: sqlite.INDEX_CONSTRAINT_EQ, OmitCheck: true}}},
	{Name: "author_login", Type: "TEXT"},
	{Name: "body", Type: "TEXT"},
	{Name: "closed", Type: "BOOLEAN"},
	{Name: "closed_at", Type: "DATETIME"},
	{Name: "comment_count", Type: "INT", OrderBy: vtab.ASC | vtab.DESC},
	{Name: "created_at", Type: "DATETIME", OrderBy: vtab.ASC | vtab.DESC},
	{Name: "created_via_email", Type: "BOOLEAN"},
	{Name: "database_id", Type: "TEXT"},
	{Name: "editor_login", Type: "TEXT"},
	{Name: "includes_created_edit", Type: "BOOLEAN"},
	{Name: "label_count", Type: "INT"},
	{Name: "labels", Type: "JSON"},
	{Name: "last_edited_at", Type: "DATETIME"},
	{Name: "locked", Type: "BOOLEAN"},
	{Name: "milestone_count", Type: "INT"},
	{Name: "number", Type: "INT"},
	{Name: "participant_count", Type: "INT"},
	{Name: "published_at", Type: "DATETIME"},
	{Name: "reaction_count", Type: "INT"},
	{Name: "state", Type: "TEXT"},
	{Name: "title", Type: "TEXT"},
	{Name: "updated_at", Type: "DATETIME", OrderBy: vtab.ASC | vtab.DESC},
	{Name: "url", Type: "TEXT"},
}

func NewIssuesModule(opts *Options) sqlite.Module {
	_ = "STUB: not implemented"
	return *new(sqlite.Module)
}
