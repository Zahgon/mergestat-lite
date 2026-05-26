package github

import (
	"context"

	"github.com/augmentable-dev/vtab"
	"github.com/mergestat/mergestat-lite/extensions/options"
	"github.com/rs/zerolog"
	"github.com/shurcooL/githubv4"
	"go.riyazali.net/sqlite"
)

type pullRequestForReviews struct {
	Id      githubv4.String
	Number  int
	Reviews struct {
		Nodes    []*prReview
		PageInfo struct {
			EndCursor   githubv4.String
			HasNextPage bool
		}
	} `graphql:"reviews(first: $perPage, after: $reviewCursor)"`
}

type prReview struct {
	Author struct {
		Login string
		Url   string
	}
	AuthorAssociation         string
	AuthorCanPushToRepository bool
	Body                      string
	Comments                  struct {
		TotalCount int
	}
	CreatedAt       githubv4.DateTime
	CreatedViaEmail bool
	Editor          struct {
		Login string
	}
	Id           string
	LastEditedAt githubv4.DateTime
	PublishedAt  githubv4.DateTime
	State        string
	SubmittedAt  githubv4.DateTime
	UpdatedAt    githubv4.DateTime
}

type fetchPRReviewsResults struct {
	RateLimit   *options.GitHubRateLimitResponse
	PullRequest *pullRequestForReviews
	HasNextPage bool
	EndCursor   *githubv4.String
}

func (i *iterPRReviews) fetchPRReviews(ctx context.Context, endCursor *githubv4.String) (*fetchPRReviewsResults, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type iterPRReviews struct {
	*Options
	owner         string
	name          string
	prNumber      int
	currentReview int
	results       *fetchPRReviewsResults
}

func (i *iterPRReviews) logger() *zerolog.Logger { _ = "STUB: not implemented"; return nil }

func (i *iterPRReviews) Column(ctx vtab.Context, c int) error {
	_ = "STUB: not implemented"
	return nil
}

func (i *iterPRReviews) Next() (vtab.Row, error) {
	_ = "STUB: not implemented"
	return *new(vtab.Row), nil
}

var prReviewCols = []vtab.Column{
	{Name: "owner", Type: "TEXT", NotNull: true, Hidden: true, Filters: []*vtab.ColumnFilter{{Op: sqlite.INDEX_CONSTRAINT_EQ, OmitCheck: true}}},
	{Name: "reponame", Type: "TEXT", NotNull: true, Hidden: true, Filters: []*vtab.ColumnFilter{{Op: sqlite.INDEX_CONSTRAINT_EQ, OmitCheck: true}}},
	{Name: "pr_number", Type: "INT", NotNull: true, Hidden: true, Filters: []*vtab.ColumnFilter{{Op: sqlite.INDEX_CONSTRAINT_EQ, OmitCheck: true}}},
	{Name: "author_login", Type: "TEXT"},
	{Name: "author_url", Type: "TEXT"},
	{Name: "author_association", Type: "TEXT"},
	{Name: "author_can_push_to_repository", Type: "BOOLEAN"},
	{Name: "body", Type: "TEXT"},
	{Name: "comment_count", Type: "INT"},
	{Name: "created_at", Type: "DATETIME"},
	{Name: "created_via_email", Type: "BOOLEAN"},
	{Name: "editor_login", Type: "TEXT"},
	{Name: "id", Type: "TEXT"},
	{Name: "last_edited_at", Type: "DATETIME"},
	{Name: "published_at", Type: "DATETIME"},
	{Name: "state", Type: "TEXT"},
	{Name: "submitted_at", Type: "DATETIME"},
	{Name: "updated_at", Type: "DATETIME"},
}

func NewPRReviewsModule(opts *Options) sqlite.Module {
	_ = "STUB: not implemented"
	return *new(sqlite.Module)
}

// if true, user supplied 3 args, 1st is org name, 2nd is repo name, 3rd is pr number
