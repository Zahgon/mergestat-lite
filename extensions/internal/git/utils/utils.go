package utils

import (
	"github.com/mergestat/mergestat-lite/extensions/services"
	"github.com/rs/zerolog"
)

// ModuleOptions holds common options for all git related modules
type ModuleOptions struct {
	Locator services.RepoLocator
	Context services.Context
	Logger  *zerolog.Logger
}

// GetDefaultRepoFromCtx looks up the defaultRepoPath key in the supplied context and returns it if set,
// otherwise it returns the current working directory
func GetDefaultRepoFromCtx(ctx services.Context) (repoPath string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}
