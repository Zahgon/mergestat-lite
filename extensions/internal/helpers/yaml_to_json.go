package helpers

import (
	"go.riyazali.net/sqlite"
)

// YamlToJson implements yaml_to_json sql function.
// The function signature of the equivalent sql function is:
//
//	yaml_to_json(string) string
type YamlToJson struct{}

func (y *YamlToJson) Args() int           { _ = "STUB: not implemented"; return 0 }
func (y *YamlToJson) Deterministic() bool { _ = "STUB: not implemented"; return false }

func (y *YamlToJson) Apply(context *sqlite.Context, value ...sqlite.Value) {
	_ = "STUB: not implemented"
	return
}
