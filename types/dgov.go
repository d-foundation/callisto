package types

import (
	dgovtypes "github.com/d-foundation/protocol/x/dgov/types"
)

// DGovOversightCommitteeParams
type DGovOversightCommitteeParams struct {
	dgovtypes.OversightCommitteeParams
	Height int64
}

func NewDGovOversightCommitteeParams(params dgovtypes.OversightCommitteeParams, height int64) DGovOversightCommitteeParams {
	return DGovOversightCommitteeParams{
		OversightCommitteeParams: params,
		Height:                   height,
	}
}

type DGovOversightCommitteeAddress struct {
	Address string
	Height  int64
}

// DGovOversightCommitteeParams
type DGovOversightCommitteeDisallowList struct {
	DisallowList []string
	Height       int64
}

func NewDGovOversightCommitteeDisallowList(list []string, height int64) DGovOversightCommitteeDisallowList {
	return DGovOversightCommitteeDisallowList{
		DisallowList: list,
		Height:       height,
	}
}
