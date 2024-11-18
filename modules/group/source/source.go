package source

import (
	group "github.com/cosmos/cosmos-sdk/x/group"
)

type Source interface {
	GetGroupInfo(id uint64) (*group.GroupInfo, error)
	GetGroupPolicyInfo(address string) (*group.GroupPolicyInfo, error)
	GetGroupMembers(id uint64) ([]*group.GroupMember, error)
}
