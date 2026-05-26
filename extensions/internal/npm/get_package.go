package npm

import (
	"go.riyazali.net/sqlite"
)

type GetPackage struct{ *Client }

func (f *GetPackage) Args() int           { _ = "STUB: not implemented"; return 0 }
func (f *GetPackage) Deterministic() bool { _ = "STUB: not implemented"; return false }
func (f *GetPackage) Apply(ctx *sqlite.Context, values ...sqlite.Value) {
	_ = "STUB: not implemented"
	return
}
