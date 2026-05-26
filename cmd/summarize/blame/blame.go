package blame

import (
	"database/sql"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/jmoiron/sqlx"
)

const preloadBlameSQL = `
CREATE TABLE preloaded_blame AS
SELECT
    files.path,
    blame.line_no,
    commits.hash,
	commits.author_name,
	commits.author_email,
	commits.author_when,
	commits.committer_name,
	commits.committer_email,
	commits.committer_when
FROM files, blame('', '', files.path)
JOIN commits ON commits.hash = blame.commit_hash
WHERE path LIKE ?
`

const blameSummarySQL = `
SELECT
	count(*) AS loc,
	count(distinct path) AS files,
	count(distinct(author_email || author_name)) AS authors,
	MAX(author_when) AS latest,
	MIN(author_when) AS oldest,
	AVG(julianday('now') - julianday(author_when)) AS avg_age,
	count(distinct hash) AS commits
FROM preloaded_blame
`

type BlameSummary struct {
	Lines   int             `db:"loc"`
	Files   int             `db:"files"`
	Authors int             `db:"authors"`
	Latest  sql.NullString  `db:"latest"`
	Oldest  sql.NullString  `db:"oldest"`
	AvgAge  sql.NullFloat64 `db:"avg_age"`
	Commits int             `db:"commits"`
}

const blameAuthorSummarySQL = `
SELECT
	author_name, author_email,
	count(*) AS loc,
	MAX(author_when) AS latest,
	MIN(author_when) AS oldest,
	AVG(julianday('now') - julianday(author_when)) AS avg_age,
	count(distinct hash) AS commits,
	json_group_array(path) AS files
FROM preloaded_blame
GROUP BY author_name, author_email
ORDER BY loc DESC
`

type BlameAuthorSummary struct {
	AuthorName  string          `db:"author_name"`
	AuthorEmail string          `db:"author_email"`
	Lines       int             `db:"loc"`
	Latest      sql.NullString  `db:"latest"`
	Oldest      sql.NullString  `db:"oldest"`
	AvgAge      sql.NullFloat64 `db:"avg_age"`
	Commits     int             `db:"commits"`
	Files       string          `db:"files"`
}

type TermUI struct {
	db                   *sqlx.DB
	pathPattern          string
	err                  error
	spinner              spinner.Model
	blamePreloaded       bool
	blameSummary         *BlameSummary
	blameAuthorSummaries *[]*BlameAuthorSummary
}

func NewTermUI(pathPattern string) (*TermUI, error) { _ = "STUB: not implemented"; return nil, nil }

func (t *TermUI) Init() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (t *TermUI) preloadBlame() tea.Msg { _ = "STUB: not implemented"; return *new(tea.Msg) }

func (t *TermUI) loadBlameSummary() tea.Msg { _ = "STUB: not implemented"; return *new(tea.Msg) }

func (t *TermUI) loadBlameAuthorSummary() tea.Msg { _ = "STUB: not implemented"; return *new(tea.Msg) }

func (t *TermUI) renderDurationString(d time.Duration) string { _ = "STUB: not implemented"; return "" }

func (t *TermUI) renderBlameSummaryTable(boldHeader bool) string {
	_ = "STUB: not implemented"
	return ""
}

func (t *TermUI) renderBlameAuthorSummary(limit int) string { _ = "STUB: not implemented"; return "" }

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
