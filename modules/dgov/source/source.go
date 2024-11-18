package source

import (
	gov "github.com/d-foundation/protocol/x/dgov/types"
)

type Source interface {
	GetOversightCommitteeParams() (*gov.OversightCommitteeParams, error)
	GetOversightCommitteeAddress() (string, error)
}
