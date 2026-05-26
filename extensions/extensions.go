// Package extensions provide implementation of the various underlying sqlite3 virtual tables [https://www.sqlite.org/vtab.html] and user defined functions
// that mergestat uses under-the-hood. This module can be side-effect-imported in other modules to include the functionality
// of the sqlite3 extensions there.
package extensions

import (
	"github.com/mergestat/mergestat-lite/extensions/options"
	"go.riyazali.net/sqlite"
)

func RegisterFn(fns ...options.OptionFn) func(ext *sqlite.ExtensionApi) (_ sqlite.ErrorCode, err error) {
	_ = "STUB: not implemented"
	return nil
}

// return an extension function that register modules with sqlite when this package is loaded

// register the git tables

// only conditionally register the utility functions

// conditionally register the GitHub functionality
