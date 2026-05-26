package github

import (
	"context"

	"github.com/augmentable-dev/vtab"
	"github.com/rs/zerolog"
	"github.com/shurcooL/githubv4"
	"go.riyazali.net/sqlite"
)

type fetchOrgAuditLogResults struct {
	AuditLogs   []*auditLogEntry
	HasNextPage bool
	EndCursor   *githubv4.String
}

type auditLogEntry struct {
	Typename     string `graphql:"__typename"`
	NodeFragment struct {
		Id string
	} `graphql:"... on Node"`
	Entry auditLogEntryContents `graphql:"... on AuditEntry"`
}

type auditLogEntryContents struct {
	Action string
	Actor  struct {
		Type string `graphql:"__typename"`
	}
	ActorLogin    string
	ActorIp       string
	ActorLocation struct {
		City        string `json:"city"`
		Country     string `json:"country"`
		CountryCode string `json:"countryCode"`
		Region      string `json:"region"`
		RegionCode  string `json:"regionCode"`
	}
	CreatedAt     githubv4.DateTime
	OperationType string
	UserLogin     string
}

func (i *iterOrgAuditLogs) fetchOrgAuditRepos(ctx context.Context, startCursor *githubv4.String) (*fetchOrgAuditLogResults, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type iterOrgAuditLogs struct {
	*Options
	login        string
	affiliations string
	current      int
	results      *fetchOrgAuditLogResults
	auditOrder   *githubv4.AuditLogOrder
}

func (i *iterOrgAuditLogs) logger() *zerolog.Logger { _ = "STUB: not implemented"; return nil }

func (i *iterOrgAuditLogs) Column(ctx vtab.Context, c int) error {
	_ = "STUB: not implemented"
	return nil
}

func (i *iterOrgAuditLogs) Next() (vtab.Row, error) {
	_ = "STUB: not implemented"
	return *new(vtab.Row), nil
}

var orgAuditCols = []vtab.Column{
	{Name: "login", Type: "TEXT", Hidden: true, Filters: []*vtab.ColumnFilter{{Op: sqlite.INDEX_CONSTRAINT_EQ, OmitCheck: true}}},
	{Name: "id", Type: "TEXT"},
	{Name: "entry_type", Type: "TEXT"},
	{Name: "action", Type: "TEXT"},
	{Name: "actor_type", Type: "TEXT"},
	{Name: "actor_login", Type: "TEXT"},
	{Name: "actor_ip", Type: "TEXT"},
	{Name: "actor_location", Type: "JSON"},
	{Name: "created_at", Type: "DATETIME", OrderBy: vtab.ASC | vtab.DESC},
	{Name: "operation_type", Type: "TEXT"},
	{Name: "user_login", Type: "TEXT"},
}

func NewOrgAuditModule(opts *Options) sqlite.Module {
	_ = "STUB: not implemented"
	return *new(sqlite.Module)
}

// for now we can only support single field order bys
