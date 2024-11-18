package remote

import (
	group "github.com/cosmos/cosmos-sdk/x/group"
	"github.com/forbole/juno/v6/node/remote"

	groupsource "github.com/forbole/callisto/v4/modules/group/source"
)

var (
	_ groupsource.Source = &Source{}
)

// Source implements stakingsource.Source using a remote node
type Source struct {
	*remote.Source
	groupClient group.QueryClient
}

// NewSource returns a new Source instance
func NewSource(source *remote.Source, groupClient group.QueryClient) *Source {
	return &Source{
		Source:      source,
		groupClient: groupClient,
	}
}

func (s *Source) GetGroupInfo(id uint64) (*group.GroupInfo, error) {
	res, err := s.groupClient.GroupInfo(s.Ctx, &group.QueryGroupInfoRequest{GroupId: id})
	if err != nil {
		return &group.GroupInfo{}, err
	}

	return res.Info, nil
}

func (s *Source) GetGroupPolicyInfo(address string) (*group.GroupPolicyInfo, error) {
	res, err := s.groupClient.GroupPolicyInfo(s.Ctx, &group.QueryGroupPolicyInfoRequest{Address: address})
	if err != nil {
		return &group.GroupPolicyInfo{}, err
	}
	return res.Info, nil
}

func (s *Source) GetGroupMembers(id uint64) ([]*group.GroupMember, error) {
	res, err := s.groupClient.GroupMembers(s.Ctx, &group.QueryGroupMembersRequest{GroupId: id})
	if err != nil {
		return []*group.GroupMember{}, err
	}
	return res.Members, nil
}
