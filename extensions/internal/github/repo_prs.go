package github

import (
	"context"

	"github.com/augmentable-dev/vtab"
	"github.com/mergestat/mergestat-lite/extensions/options"
	"github.com/rs/zerolog"
	"github.com/shurcooL/githubv4"
	"go.riyazali.net/sqlite"
)

type pullRequest struct {
	ActiveLockReason githubv4.LockReason
	Additions        int
	Author           struct {
		Login     string
		AvatarUrl *githubv4.URI
		User      struct {
			Name string
		} `graphql:"... on User"`
	}
	AuthorAssociation githubv4.CommentAuthorAssociation
	BaseRefOid        githubv4.GitObjectID
	BaseRefName       string
	BaseRepository    struct {
		NameWithOwner string
	}
	Body         string
	ChangedFiles int
	Closed       bool
	ClosedAt     githubv4.DateTime
	Comments     struct {
		TotalCount int
	}
	Commits struct {
		TotalCount int
	}
	CreatedAt       githubv4.DateTime
	CreatedViaEmail bool
	DatabaseID      int
	Deletions       int
	Editor          struct {
		Login string
	}
	HeadRefName    string
	HeadRefOid     githubv4.GitObjectID
	HeadRepository struct {
		NameWithOwner string
	}
	IsCrossRepository bool
	IsDraft           bool
	Labels            struct {
		TotalCount int
		Nodes      []struct {
			Name string
		}
	} `graphql:"labels(first: 15)"`
	LastEditedAt        githubv4.DateTime
	Locked              bool
	MaintainerCanModify bool
	Mergeable           githubv4.MergeableState
	Merged              bool
	MergedAt            githubv4.DateTime
	MergedBy            struct {
		Login string
	}
	Number       int
	Participants struct {
		TotalCount int
	}
	PublishedAt    githubv4.DateTime
	ReviewDecision githubv4.PullRequestReviewDecision
	State          githubv4.PullRequestState
	Title          string
	UpdatedAt      githubv4.DateTime
	Url            githubv4.URI
}

type fetchPRResults struct {
	RateLimit   *options.GitHubRateLimitResponse
	Edges       []*pullRequest
	HasNextPage bool
	EndCursor   *githubv4.String
}

func (i *iterPRs) fetchPRs(ctx context.Context, startCursor *githubv4.String) (*fetchPRResults, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type iterPRs struct {
	*Options
	owner   string
	name    string
	current int
	results *fetchPRResults
	prOrder *githubv4.IssueOrder
}

func (i *iterPRs) logger() *zerolog.Logger { _ = "STUB: not implemented"; return nil }

func (i *iterPRs) Column(ctx vtab.Context, c int) error { _ = "STUB: not implemented"; return nil }

func (i *iterPRs) Next() (vtab.Row, error) { _ = "STUB: not implemented"; return *new(vtab.Row), nil }

var prCols = []vtab.Column{
	{Name: "owner", Type: "TEXT", NotNull: true, Hidden: true, Filters: []*vtab.ColumnFilter{{Op: sqlite.INDEX_CONSTRAINT_EQ, OmitCheck: true}}},
	{Name: "reponame", Type: "TEXT", NotNull: true, Hidden: true, Filters: []*vtab.ColumnFilter{{Op: sqlite.INDEX_CONSTRAINT_EQ, OmitCheck: true}}},
	{Name: "additions", Type: "INT"},
	{Name: "author_login", Type: "TEXT"},
	{Name: "author_association", Type: "TEXT"},
	{Name: "author_avatar_url", Type: "TEXT"},
	{Name: "author_name", Type: "TEXT"},
	{Name: "base_ref_oid", Type: "TEXT"},
	{Name: "base_ref_name", Type: "TEXT"},
	{Name: "base_repository_name", Type: "TEXT"},
	{Name: "body", Type: "TEXT"},
	{Name: "changed_files", Type: "INT"},
	{Name: "closed", Type: "BOOLEAN"},
	{Name: "closed_at", Type: "DATETIME"},
	{Name: "comment_count", Type: "INT", OrderBy: vtab.ASC | vtab.DESC, Filters: []*vtab.ColumnFilter{
		{Op: sqlite.INDEX_CONSTRAINT_GT}, {Op: sqlite.INDEX_CONSTRAINT_GE},
		{Op: sqlite.INDEX_CONSTRAINT_LT}, {Op: sqlite.INDEX_CONSTRAINT_LE},
	}},
	{Name: "commit_count", Type: "INT"},
	{Name: "created_at", Type: "DATETIME", OrderBy: vtab.ASC | vtab.DESC, Filters: []*vtab.ColumnFilter{
		{Op: sqlite.INDEX_CONSTRAINT_GT}, {Op: sqlite.INDEX_CONSTRAINT_GE},
		{Op: sqlite.INDEX_CONSTRAINT_LT}, {Op: sqlite.INDEX_CONSTRAINT_LE},
	}},
	{Name: "created_via_email", Type: "BOOLEAN"},
	{Name: "database_id", Type: "INT"},
	{Name: "deletions", Type: "INT"},
	{Name: "editor_login", Type: "TEXT"},
	{Name: "head_ref_name", Type: "TEXT"},
	{Name: "head_ref_oid", Type: "TEXT"},
	{Name: "head_repository_name", Type: "TEXT"},
	{Name: "is_draft", Type: "BOOLEAN"},
	{Name: "label_count", Type: "INT"},
	{Name: "labels", Type: "JSON"},
	{Name: "last_edited_at", Type: "DATETIME"},
	{Name: "locked", Type: "BOOLEAN"},
	{Name: "maintainer_can_modify", Type: "BOOLEAN"},
	{Name: "mergeable", Type: "TEXT"},
	{Name: "merged", Type: "BOOLEAN"},
	{Name: "merged_at", Type: "DATETIME"},
	{Name: "merged_by", Type: "TEXT"},
	{Name: "number", Type: "INT"},
	{Name: "participant_count", Type: "INT"},
	{Name: "published_at", Type: "DATETIME"},
	{Name: "review_decision", Type: "TEXT"},
	{Name: "state", Type: "TEXT"},
	{Name: "title", Type: "TEXT"},
	{Name: "updated_at", Type: "DATETIME", OrderBy: vtab.ASC | vtab.DESC, Filters: []*vtab.ColumnFilter{
		{Op: sqlite.INDEX_CONSTRAINT_GT}, {Op: sqlite.INDEX_CONSTRAINT_GE},
		{Op: sqlite.INDEX_CONSTRAINT_LT}, {Op: sqlite.INDEX_CONSTRAINT_LE},
	}},
	{Name: "url", Type: "TEXT"},
}

func NewPRModule(opts *Options) sqlite.Module {
	_ = "STUB: not implemented"
	return *new(sqlite.Module)
}
