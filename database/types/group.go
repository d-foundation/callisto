package types

import "time"

// Group Row

// GroupRow represents a single row inside the groups table
type GroupRow struct {
	ID           int64     `db:"id"`
	AdminAddress string    `db:"admin_address"`
	Metadata     string    `db:"metadata"`
	Version      int64     `db:"version"`
	CreatedAt    time.Time `db:"created_at"`
}

// GroupMembersRow represents a single row inside the group_members table
type GroupMembersRow struct {
	GroupID       int64     `db:"group_id"`
	MemberAddress string    `db:"member_address"`
	Weight        string    `db:"weight"`
	Metadata      string    `db:"metadata"`
	AddedAt       time.Time `db:"added_at"`
}

// GroupPoliciesRow represents a single row inside the group_policies table
type GroupPoliciesRow struct {
	Address        string    `db:"address"`
	GroupID        int64     `db:"group_id"`
	AdminAddress   string    `db:"admin_address"`
	Metadata       string    `db:"metadata"`
	Version        int64     `db:"version"`
	DecisionPolicy string    `db:"decision_policy"`
	CreatedAt      time.Time `db:"created_at"`
}
