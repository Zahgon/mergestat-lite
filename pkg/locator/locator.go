// Package locator provides various different implementations of
// git.RepoLocator service. The locator service is used by the implementations
// of Git virtual modules to query for / locate a given repository.
//
// The various different implementations of this interface provided in this package
// provide different ways to locate the service, while some provides additional services
// such as caching or switching between implementations.
package locator

import (
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/mergestat/mergestat-lite/extensions/services"
	"github.com/rs/zerolog"
)

// DiskLocator is a repo locator implementation that opens on-disk repository at the specified path.
func DiskLocator() services.RepoLocator {
	_ = "STUB: not implemented"
	return *new(services.RepoLocator)
}

// CachedLocator is decorator function that takes a RepoLocator instance
// and returns another one that caches output from the underlying locator
// using path as the key.
func CachedLocator(rl services.RepoLocator) services.RepoLocator {
	_ = "STUB: not implemented"
	return *new(services.RepoLocator)
}

// determineCloneDir returns the path to a directory on disk where a repository will be cloned to
// given a baseCloneDir. If baseCloneDir == "", a tmp dir will be created, otherwise a directory
// path will be determined based on the URL (HTTP(s) or SSH) of the provided repository.
// The bool returned (2nd return val) indicates whether the output dir is in a tmp directory or not.
func determineCloneDir(path, baseCloneDir string) (string, bool, error) {
	_ = "STUB: not implemented"
	return "", false, nil
}

// if no clone directory is specified, use a tmp dir

// if clone directory is set, get the abs path

// then use the parsed path to determine where repos should end up

// assume it's an ssh repo

// HttpLocator returns a repo locator capable of cloning remote
// http repositories on-demand into temporary storage. It is recommended
// that you club it with something like CachedLocator to improve performance
// and remove the need to clone a single repository multiple times.
func HttpLocator(o *MultiLocatorOptions) func() services.RepoLocator {
	_ = "STUB: not implemented"
	return nil
}

// httpLocatorWithAuth returns a func that returns a repo locator 🤯
// its primary intended use is below in the MultiLocator, which receives options.
// If HTTP auth options are supplied, they will be used when cloning an https (only https) repo.
func httpLocatorWithAuth(user, pass string, rl services.RepoLocator) func() services.RepoLocator {
	_ = "STUB: not implemented"
	return nil
}

// SSHLocator returns a repo locator capable of cloning remote
// ssh repositories on-demand into temporary storage. It is recommended
// that you club it with something like CachedLocator to improve performance
// and remove the need to clone a single repository multiple times.
func SSHLocator(o *MultiLocatorOptions) func() services.RepoLocator {
	_ = "STUB: not implemented"
	return nil
}

// TODO(patrickdevivo) maybe a little hacky instead of properly parsing the url, strip out the username first
// if it's set, otherwise default to "git"

type MultiLocatorOptions struct {
	HTTPAuth        *http.BasicAuth
	CloneDir        string
	InsecureSkipTLS bool
}

// MultiLocator returns a locator service that work with multiple git protocols
// and is able to pick the correct underlying locator based on path provided.
func MultiLocator(o *MultiLocatorOptions) services.RepoLocator {
	_ = "STUB: not implemented"
	return *new(services.RepoLocator)
}

// file is the default locator

// LoggingLocator returns a locator that logs
func LoggingLocator(logger *zerolog.Logger, rl services.RepoLocator) services.RepoLocator {
	_ = "STUB: not implemented"
	return *new(services.RepoLocator)
}
