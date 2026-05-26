package helpers

import (
	"go.riyazali.net/sqlite"
)

// TomlToJson implements toml_to_json sql function.
// The function signature of the equivalent sql function is:
//
//	toml_to_json(string) string
type TomlToJson struct{}

func (y *TomlToJson) Args() int           { _ = "STUB: not implemented"; return 0 }
func (y *TomlToJson) Deterministic() bool { _ = "STUB: not implemented"; return false }

func (y *TomlToJson) Apply(context *sqlite.Context, value ...sqlite.Value) {
	_ = "STUB: not implemented"
	return
}
