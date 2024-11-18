package remote

import (
	dgov "github.com/d-foundation/protocol/x/dgov/types"
	"github.com/forbole/juno/v6/node/remote"

	dgovsource "github.com/forbole/callisto/v4/modules/dgov/source"
)

var (
	_ dgovsource.Source = &Source{}
)

// Source implements stakingsource.Source using a remote node
type Source struct {
	*remote.Source
	dgovClient dgov.QueryClient
}

// NewSource returns a new Source instance
func NewSource(source *remote.Source, client dgov.QueryClient) *Source {
	return &Source{
		Source:     source,
		dgovClient: client,
	}
}

func (s *Source) GetOversightCommitteeParams() (*dgov.OversightCommitteeParams, error) {
	res, err := s.dgovClient.GetOversightCommitteeParams(s.Ctx, &dgov.GetOversightCommitteeParamsRequest{})
	if err != nil {
		return &dgov.OversightCommitteeParams{}, err
	}
	return res.Params, nil
}

func (s *Source) GetOversightCommitteeAddress() (string, error) {
	res, err := s.dgovClient.GetOversightCommitteeAddr(s.Ctx, &dgov.GetOversightCommitteeAddrRequest{})
	if err != nil {
		return "", err
	}
	return res.OversightCommitteeAddr, nil
}
