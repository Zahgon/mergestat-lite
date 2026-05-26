package cmd

import (

	// bring in sqlite 🙌
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/mergestat/mergestat-lite/pkg/sqlite"
)

func registerExt() { _ = "STUB: not implemented"; return }
