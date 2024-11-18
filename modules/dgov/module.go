package dgov

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/forbole/juno/v6/modules"

	"github.com/forbole/callisto/v4/database"
	dgovsource "github.com/forbole/callisto/v4/modules/dgov/source"
)

var (
	_ modules.Module        = &Module{}
	_ modules.GenesisModule = &Module{}
)

// Module represents the x/group module
type Module struct {
	cdc    codec.Codec
	db     *database.Db
	source dgovsource.Source
}

// NewModule returns a new Module instance
func NewModule(
	source dgovsource.Source, cdc codec.Codec, db *database.Db,
) *Module {
	return &Module{
		cdc:    cdc,
		db:     db,
		source: source,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "dgov"
}
