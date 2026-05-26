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

type fetchOrgReposResults struct {
	RateLimit   *options.GitHubRateLimitResponse
	OrgRepos    []*orgRepo
	HasNextPage bool
	EndCursor   *githubv4.String
}

type orgRepo struct {
	CreatedAt        time.Time
	DatabaseId       int
	DefaultBranchRef struct {
		Name   string
		Prefix string
	}
	Description string
	DiskUsage   int
	ForkCount   int
	HomepageUrl string
	IsArchived  bool
	IsDisabled  bool
	IsFork      bool
	IsMirror    bool
	IsPrivate   bool
	Issues      struct {
		TotalCount int
	}
	LatestRelease struct {
		Author struct {
			Login string
		}
		CreatedAt   githubv4.DateTime
		Name        string
		PublishedAt githubv4.DateTime
	}
	LicenseInfo struct {
		Key      string
		Name     string
		Nickname string
	}
	Name              string
	OpenGraphImageUrl githubv4.URI
	PrimaryLanguage   struct {
		Name string
	}
	PullRequests struct {
		TotalCount int
	}
	PushedAt time.Time
	Releases struct {
		TotalCount int
	}
	StargazerCount int
	Topics         struct {
		Nodes []struct {
			Topic struct {
				Name string
			}
		}
	} `graphql:"repositoryTopics(first: 10)"`
	UpdatedAt time.Time
	Watchers  struct {
		TotalCount int
	}
}

func (i *iterOrgRepos) fetchOrgRepos(ctx context.Context, startCursor *githubv4.String) (*fetchOrgReposResults, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type iterOrgRepos struct {
	*Options
	login        string
	affiliations string
	current      int
	results      *fetchOrgReposResults
	repoOrder    *githubv4.RepositoryOrder
}

func (i *iterOrgRepos) logger() *zerolog.Logger { _ = "STUB: not implemented"; return nil }

func (i *iterOrgRepos) Column(ctx vtab.Context, c int) error { _ = "STUB: not implemented"; return nil }

func (i *iterOrgRepos) Next() (vtab.Row, error) {
	_ = "STUB: not implemented"
	return *new(vtab.Row), nil
}

var orgReposCols = []vtab.Column{
	{Name: "login", Type: "TEXT", Hidden: true, Filters: []*vtab.ColumnFilter{{Op: sqlite.INDEX_CONSTRAINT_EQ, OmitCheck: true}}},
	{Name: "affiliations", Type: "TEXT", Hidden: true, Filters: []*vtab.ColumnFilter{{Op: sqlite.INDEX_CONSTRAINT_EQ, OmitCheck: true}}},
	{Name: "created_at", Type: "DATETIME", OrderBy: vtab.ASC | vtab.DESC},
	{Name: "database_id", Type: "INT"},
	{Name: "default_branch_ref_name", Type: "TEXT"},
	{Name: "default_branch_ref_prefix", Type: "TEXT"},
	{Name: "description", Type: "TEXT"},
	{Name: "disk_usage", Type: "INT"},
	{Name: "fork_count", Type: "INT"},
	{Name: "homepage_url", Type: "TEXT"},
	{Name: "is_archived", Type: "BOOLEAN"},
	{Name: "is_disabled", Type: "BOOLEAN"},
	{Name: "is_fork", Type: "BOOLEAN"},
	{Name: "is_mirror", Type: "BOOLEAN"},
	{Name: "is_private", Type: "BOOLEAN"},
	{Name: "issue_count", Type: "INT"},
	{Name: "latest_release_author", Type: "TEXT"},
	{Name: "latest_release_created_at", Type: "DATETIME"},
	{Name: "latest_release_name", Type: "TEXT"},
	{Name: "latest_release_published_at", Type: "DATETIME"},
	{Name: "license_key", Type: "TEXT"},
	{Name: "license_name", Type: "TEXT"},
	{Name: "name", Type: "TEXT", OrderBy: vtab.ASC | vtab.DESC},
	{Name: "open_graph_image_url", Type: "TEXT"},
	{Name: "primary_language", Type: "TEXT"},
	{Name: "pull_request_count", Type: "INT"},
	{Name: "pushed_at", Type: "DATETIME", OrderBy: vtab.ASC | vtab.DESC},
	{Name: "release_count", Type: "INT"},
	{Name: "stargazer_count", Type: "INT", OrderBy: vtab.ASC | vtab.DESC},
	{Name: "topics", Type: "JSON"},
	{Name: "updated_at", Type: "DATETIME", OrderBy: vtab.ASC | vtab.DESC},
	{Name: "watcher_count", Type: "INT"},
}

func NewOrgReposModule(opts *Options) sqlite.Module {
	_ = "STUB: not implemented"
	return *new(sqlite.Module)
}

// for now we can only support single field order bys
