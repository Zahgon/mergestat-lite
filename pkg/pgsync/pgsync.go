package pgsync

import (
	"context"
	"database/sql"

	"github.com/rs/zerolog"

	_ "github.com/mattn/go-sqlite3"
	_ "github.com/mergestat/mergestat-lite/pkg/sqlite"
)

type SyncOptions struct {
	Postgres   *sql.DB
	MergeStat  *sql.DB
	SchemaName string
	TableName  string
	Query      string
	Logger     *zerolog.Logger
}

// Sync imports the results of an mergestat query into a postgres table.
// CAUTION: will overwrite (DROP!) the specified table and replace it.
func Sync(ctx context.Context, options *SyncOptions) error { _ = "STUB: not implemented"; return nil }

// create a new temp table

// sqliteTypeToPostgresType maps SQLite column types to Postgres column types
func sqliteTypeToPostgresType(col *sql.ColumnType) string {
	_ = "STUB: not implemented"
	// TODO(patrickdevivo) expressions do not have a type-affinity in SQLite (unless explicitly cast)
	// which means something like `datetime('now')` will not have a type-affinity and be rendered into postgres
	// as text. Even `CAST(datetime('now') AS "DATETIME")` won't work because "DATETIME" is not a known affinity (becomes numeric).
	// All this means that there's not a way, using a SQL query alone, to coerce a column into a specific postgres type.
	// This matters mainly when there are expressions in a query (column name references will use the declared column type)
	return ""
}

// createTableFromSQLiteTypes produces a postgres CREATE TABLE statement from a set of SQLite columns
func createTableFromSQLiteTypes(schemaName, tableName string, columns []*sql.ColumnType) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// helper to determine whether we're on the last column (and therefore should avoid a comma ",") in the range
