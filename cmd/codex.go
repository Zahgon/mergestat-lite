package cmd

import (
	"context"
	_ "embed"
)

//go:embed codex-prompt-context.sql
var promptPrefix string

// codexToSQL generates SQL from a natural language prompt
func codexToSQL(ctx context.Context, prompt string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
