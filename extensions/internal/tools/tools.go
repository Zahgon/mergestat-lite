// Package tools provides utilities to help with writing
// integration test suites for mergestat sqlite module.
package tools

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	_ "github.com/mergestat/mergestat-lite/pkg/sqlite"
)

func RowContent(rows *sql.Rows) (colCount int, contents [][]string, err error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}
