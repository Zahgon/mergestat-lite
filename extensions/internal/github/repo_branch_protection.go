package github

import (
	"context"

	"github.com/augmentable-dev/vtab"
	"github.com/mergestat/mergestat-lite/extensions/options"
	"github.com/rs/zerolog"
	"github.com/shurcooL/githubv4"
	"go.riyazali.net/sqlite"
)

/* TODO: Consider adding
* BranchProtectionRuleConflicts
* MatchingRefs
* PushAllowances
* ReviewDismissalAllowances
 */
type protectionRules struct {
	AllowsDeletions   bool
	AllowsForcePushes bool
	Creator           struct {
		Login string
	}
	DatabaseId                     int
	DismissesStaleReviews          bool
	IsAdminEnforced                bool
	Pattern                        string
	RequiredApprovingReviewCount   int
	RequiredStatusCheckContexts    []string
	RequiresApprovingReviews       bool
	RequiresCodeOwnerReviews       bool
	RequiresCommitSignatures       bool
	RequiresConversationResolution bool
	RequiresLinearHistory          bool
	RequiresStatusChecks           bool
	RequiresStrictStatusChecks     bool
	RestrictsPushes                bool
	RestrictsReviewDismissals      bool
}

type fetchBranchProtectionResults struct {
	RateLimit   *options.GitHubRateLimitResponse
	Edges       []*protectionRules
	HasNextPage bool
	EndCursor   *githubv4.String
}

func (i *iterProtections) fetchProtections(ctx context.Context, startCursor *githubv4.String) (*fetchBranchProtectionResults, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type iterProtections struct {
	*Options
	owner   string
	name    string
	current int
	results *fetchBranchProtectionResults
}

func (i *iterProtections) logger() *zerolog.Logger { _ = "STUB: not implemented"; return nil }

func (i *iterProtections) Column(ctx vtab.Context, c int) error {
	_ = "STUB: not implemented"
	return nil
}

func (i *iterProtections) Next() (vtab.Row, error) {
	_ = "STUB: not implemented"
	return *new(vtab.Row), nil
}

var protectionCols = []vtab.Column{
	{Name: "owner", Type: "TEXT", NotNull: true, Hidden: true, Filters: []*vtab.ColumnFilter{{Op: sqlite.INDEX_CONSTRAINT_EQ, OmitCheck: true}}},
	{Name: "reponame", Type: "TEXT", NotNull: true, Hidden: true, Filters: []*vtab.ColumnFilter{{Op: sqlite.INDEX_CONSTRAINT_EQ, OmitCheck: true}}},
	{Name: "allow_deletions", Type: "BOOLEAN"},
	{Name: "allows_force_pushes", Type: "BOOLEAN"},
	{Name: "creator_login", Type: "TEXT"},
	{Name: "database_id", Type: "INT"},
	{Name: "dismisses_stale_reviews", Type: "BOOLEAN"},
	{Name: "is_admin_enforced", Type: "BOOLEAN"},
	{Name: "pattern", Type: "TEXT"},
	{Name: "required_approving_review_count", Type: "INT"},
	{Name: "required_status_check_contexts", Type: "BOOLEAN"},
	{Name: "requires_approving_reviews", Type: "BOOLEAN"},
	{Name: "requires_code_owners_reviews", Type: "BOOLEAN"},
	{Name: "requires_commit_signature", Type: "BOOLEAN"},
	{Name: "requires_conversation_resolution", Type: "BOOLEAN"},
	{Name: "requires_linear_history", Type: "BOOLEAN"},
	{Name: "requires_status_checks", Type: "BOOLEAN"},
	{Name: "requires_strict_status_checks", Type: "BOOLEAN"},
	{Name: "restricts_pushes", Type: "BOOLEAN"},
	{Name: "restricts_review_dismissal", Type: "BOOLEAN"},
}

func NewProtectionsModule(opts *Options) sqlite.Module {
	_ = "STUB: not implemented"
	return *new(sqlite.Module)
}
