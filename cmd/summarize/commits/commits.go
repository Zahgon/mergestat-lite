package commits

import (
	"database/sql"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/jmoiron/sqlx"
)

type CommitSummary struct {
	Total           int            `db:"total"`
	TotalNonMerges  int            `db:"total_non_merges"`
	FirstCommit     sql.NullString `db:"first_commit"`
	LastCommit      sql.NullString `db:"last_commit"`
	DistinctAuthors int            `db:"distinct_authors"`
	DistinctFiles   int            `db:"distinct_files"`
}

// This is a bit odd. We have two queries, one that includes filtering by file_path with a LIKE
// and another that skips file_path filtering completely.
// This is because it is possible to have "empty" commits - commits where no changes were made
// (and therefore there are no file_paths to filter on). For these commits, stats.* will be all NULLs.
// If a user does *not* specify a file pattern to match on, we want to include these empty commits
// in our calculations, therefore we don't mention file_path in the WHERE clause at all.
//
// This is because `file_path LIKE %` will not match when file_path IS NULL (empty commit).
// When a path filter is supplied by the user, we do apply it. Note that even supplying just a '%'
// will exclude empty commits from the resultset. This makes sense, because empty commits won't have
// changed any files in the specified pattern (they won't have changed any files at all).
const preloadCommitsWithFilePathPatternSQL = `
CREATE TABLE preloaded_commit_stats AS SELECT * FROM commits LEFT JOIN stats('', commits.hash) WHERE file_path LIKE $file_path AND author_when > date($start, $start_mod) AND author_when < date($end, $end_mod);
CREATE TABLE preloaded_commits AS SELECT hash, author_name, author_email, author_when, parents FROM preloaded_commit_stats GROUP BY hash;
`

// See comment above
const preloadCommitsWithoutFilePathPatternSQL = `
CREATE TABLE preloaded_commit_stats AS SELECT * FROM commits LEFT JOIN stats('', commits.hash) WHERE author_when > date($start, $start_mod) AND author_when < date($end, $end_mod);
CREATE TABLE preloaded_commits AS SELECT hash, author_name, author_email, author_when, parents FROM preloaded_commit_stats GROUP BY hash;
`

const commitSummarySQL = `
SELECT
	(SELECT count(*) FROM preloaded_commits) AS total,
	(SELECT count(*) FROM preloaded_commits WHERE parents < 2) AS total_non_merges,
	(SELECT author_when FROM preloaded_commits ORDER BY author_when ASC LIMIT 1) AS first_commit,
	(SELECT author_when FROM preloaded_commits ORDER BY author_when DESC LIMIT 1) AS last_commit,
	(SELECT count(distinct(author_email || author_name)) FROM preloaded_commits) AS distinct_authors,
	(SELECT count(distinct(file_path)) FROM preloaded_commit_stats WHERE file_path LIKE $file_path) AS distinct_files
`

type CommitAuthorSummary struct {
	AuthorName    string         `db:"author_name"`
	AuthorEmail   sql.NullString `db:"author_email"`
	Commits       int            `db:"commit_count"`
	Additions     sql.NullInt64  `db:"additions"`
	Deletions     sql.NullInt64  `db:"deletions"`
	DistinctFiles int            `db:"distinct_files"`
	FirstCommit   string         `db:"first_commit"`
	LastCommit    string         `db:"last_commit"`
}

const commitAuthorSummarySQL = `
SELECT
	author_name, author_email,
	count(distinct hash) AS commit_count,
	sum(additions) AS additions,
	sum(deletions) AS deletions,
	count(distinct file_path) AS distinct_files,
	min(author_when) AS first_commit,
	max(author_when) AS last_commit
FROM preloaded_commit_stats
GROUP BY author_name, author_email
ORDER BY commit_count DESC
`

type dateFilter struct {
	date string
	mod  string
}

type TermUI struct {
	db                    *sqlx.DB
	pathPattern           string
	dateFilterStart       dateFilter
	dateFilterEnd         dateFilter
	err                   error
	spinner               spinner.Model
	commitsPreloaded      bool
	commitSummary         *CommitSummary
	commitAuthorSummaries *[]*CommitAuthorSummary
}

func NewTermUI(pathPattern, dateFilterStart, dateFilterEnd string) (*TermUI, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// if the start date cannot be parsed, assume it is a date modifier relative to 'now'

func (t *TermUI) Init() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (t *TermUI) preloadCommits() tea.Msg { _ = "STUB: not implemented"; return *new(tea.Msg) }

func (t *TermUI) loadCommitSummary() tea.Msg { _ = "STUB: not implemented"; return *new(tea.Msg) }

func (t *TermUI) loadAuthorCommitSummary() tea.Msg { _ = "STUB: not implemented"; return *new(tea.Msg) }

func (t *TermUI) renderCommitSummaryTable(boldHeader bool) string {
	_ = "STUB: not implemented"
	return ""
}

func (t *TermUI) renderCommitAuthorSummary(limit int) string { _ = "STUB: not implemented"; return "" }

func (t *TermUI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(tea.Model), *new(tea.Cmd)
}

func (t *TermUI) View() string { _ = "STUB: not implemented"; return "" }

// PrintNoTTY prints a version of output with no terminal styles
func (t *TermUI) PrintNoTTY() string { _ = "STUB: not implemented"; return "" }

// PrintJSON outputs summary results as a JSON object
func (t *TermUI) PrintJSON() string { _ = "STUB: not implemented"; return "" }

func (t *TermUI) Close() error { _ = "STUB: not implemented"; return nil }
