package group

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/forbole/juno/v6/modules"

	"github.com/forbole/callisto/v4/database"
	groupsource "github.com/forbole/callisto/v4/modules/group/source"
)

var (
	_ modules.Module        = &Module{}
	_ modules.GenesisModule = &Module{}
)

// Module represents the x/group module
type Module struct {
	cdc    codec.Codec
	db     *database.Db
	source groupsource.Source
}

// NewModule returns a new Module instance
func NewModule(
	source groupsource.Source, cdc codec.Codec, db *database.Db,
) *Module {
	return &Module{
		cdc:    cdc,
		db:     db,
		source: source,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "group"
}
