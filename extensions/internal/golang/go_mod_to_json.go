package golang

import (
	"go.riyazali.net/sqlite"
	"golang.org/x/mod/module"
)

type GoModToJSON struct{}

type Version struct {
	Path    string `json:"path"`
	Version string `json:"version,omitempty"`
}

type VersionInterval struct {
	Low  string `json:"low"`
	High string `json:"high"`
}

type Require struct {
	Mod      Version `json:"mod"`
	Indirect bool    `json:"indirect,omitempty"`
}

type Exclude struct {
	Mod Version
}

type Replace struct {
	Old Version `json:"old"`
	New Version `json:"new"`
}

type Retract struct {
	VersionInterval
	Rationale string `json:"rationale"`
}

type GoModFile struct {
	Version Version    `json:"version"`
	Go      string     `json:"go"`
	Require []*Require `json:"require"`
	Exclude []*Exclude `json:"exclude"`
	Replace []*Replace `json:"replace"`
	Retract []*Retract `json:"retract"`
}

func goModVersionToVersion(v module.Version) Version {
	_ = "STUB: not implemented"
	return *new(Version)
}

func (f *GoModToJSON) Args() int           { _ = "STUB: not implemented"; return 0 }
func (f *GoModToJSON) Deterministic() bool { _ = "STUB: not implemented"; return false }
func (f *GoModToJSON) Apply(context *sqlite.Context, value ...sqlite.Value) {
	_ = "STUB: not implemented"
	return
}
