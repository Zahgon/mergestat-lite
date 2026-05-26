package display

import (
	"database/sql"
	"io"
)

func WriteTo(rows *sql.Rows, w io.Writer, format string, interactive bool) error {
	_ = "STUB: not implemented"
	return nil
}

func single(rows *sql.Rows, write io.Writer) error { _ = "STUB: not implemented"; return nil }

func csvDisplay(rows *sql.Rows, commaChar rune, noHeader bool, writer io.Writer) error {
	_ = "STUB: not implemented"
	return nil
}

func ndjsonDisplay(rows *sql.Rows, writer io.Writer) error { _ = "STUB: not implemented"; return nil }

func jsonDisplay(rows *sql.Rows, writer io.Writer) error { _ = "STUB: not implemented"; return nil }

func tableDisplay(rows *sql.Rows, write io.Writer, overflow bool) error {
	_ = "STUB: not implemented"
	return nil
}

//  TODO - getting terminal size seems to fail with `operation not supported by device` in tests
//  as a workaround for now, set a default width instead of returning an error, if one is encountered
