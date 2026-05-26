package github

import (
	"github.com/mergestat/mergestat-lite/extensions/options"
	"github.com/mergestat/mergestat-lite/extensions/services"
	"github.com/rs/zerolog"
	"github.com/shurcooL/githubv4"
	"golang.org/x/time/rate"
)

type Options struct {
	Client                func() *githubv4.Client
	RateLimiter           *rate.Limiter
	RateLimitHandler      func(*options.GitHubRateLimitResponse)
	GitHubPreRequestHook  func()
	GitHubPostRequestHook func()
	// PerPage is the default number of items per page to use when making a paginated GitHub API request
	PerPage int
	Logger  *zerolog.Logger
}

// GetGitHubTokenFromCtx looks up the githubToken key in the supplied context and returns it if set
func GetGitHubTokenFromCtx(ctx services.Context) string { _ = "STUB: not implemented"; return "" }

// GetGitHubRateLimitFromCtx looks up the githubRateLimit key in the supplied context and parses it to return a client
// side rate limit in the form "(number of reqs)/(number of seconds)". For instance a string "2/3" would yield a rate limiter
// that permis 2 requests every 3 seconds. A single integer is also permitted, which assumes the "denominator" is 1 second.
// So a value of "5" would simple mean 5 requests per second.
// If the string cannot be parsed, nil is returned.
func GetGitHubRateLimitFromCtx(ctx services.Context) *rate.Limiter {
	_ = "STUB: not implemented"
	return nil
}

// GetGitHubPerPageFromCtx looks up the githubPerPage key in the supplied context and returns it if set,
// otherwise it returns a default of 50
func GetGitHubPerPageFromCtx(ctx services.Context) int { _ = "STUB: not implemented"; return 0 }

// t1f0 converts a bool to an int
func t1f0(b bool) int { _ = "STUB: not implemented"; return 0 }

// orderByToGitHubOrder is a helper that takes a boolean indicating whether DESC or ASC and returns
// a corresponding OrderDirection from the githubv4 library
func orderByToGitHubOrder(desc bool) githubv4.OrderDirection {
	_ = "STUB: not implemented"
	return *new(githubv4.OrderDirection)
}

// repoOwnerAndName returns the "owner" and "name" (respective return values) or an error
// given the inputs to the iterator. This allows for both `SELECT * FROM github_table('mergestat/mergestat')`
// and `SELECT * FROM github_table('mergestat', 'mergestat')
func repoOwnerAndName(name, fullNameOrOwner string) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

// affiliationsFromString takes a CSV list of repository affiliations as text
// and returns a list for use in the GraphQL request
func affiliationsFromString(affiliations string) []githubv4.RepositoryAffiliation {
	_ = "STUB: not implemented"
	return nil
}
