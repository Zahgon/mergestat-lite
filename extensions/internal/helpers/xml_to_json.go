package helpers

import (
	"go.riyazali.net/sqlite"
)

// XmlToJson implements xml_to_json sql function.
// The function signature of the equivalent sql function is:
//
//	xml_to_json(string) string
type XmlToJson struct{}

func (y *XmlToJson) Args() int           { _ = "STUB: not implemented"; return 0 }
func (y *XmlToJson) Deterministic() bool { _ = "STUB: not implemented"; return false }

func (y *XmlToJson) Apply(context *sqlite.Context, value ...sqlite.Value) {
	_ = "STUB: not implemented"
	return
}
