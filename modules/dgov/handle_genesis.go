package dgov

import (
	"encoding/json"
	"fmt"

	tmtypes "github.com/cometbft/cometbft/types"

	dgovtypes "github.com/d-foundation/protocol/x/dgov/types"
	"github.com/forbole/callisto/v4/types"
	"github.com/rs/zerolog/log"
)

// HandleGenesis implements modules.Module
func (m *Module) HandleGenesis(doc *tmtypes.GenesisDoc, appState map[string]json.RawMessage) error {
	log.Debug().Str("module", "gov").Msg("parsing genesis")

	// Read the genesis state
	var genStatedgov dgovtypes.GenesisState
	err := m.cdc.UnmarshalJSON(appState[dgovtypes.ModuleName], &genStatedgov)
	if err != nil {
		return fmt.Errorf("error while reading gov genesis data: %s", err)
	}

	// Save the params
	err = m.db.SaveDGovOCParams(types.NewDGovOversightCommitteeParams(*genStatedgov.GetOversightCommitteeParams(), doc.InitialHeight))
	if err != nil {
		return fmt.Errorf("error while storing genesis dgov oc params: %s", err)
	}

	// Save the oc address
	err = m.db.SaveDGovOCAddress(types.DGovOversightCommitteeAddress{
		Address: genStatedgov.GetOversightCommitteeAddr(),
		Height:  doc.InitialHeight,
	})
	if err != nil {
		return fmt.Errorf("error while storing genesis dgov oc addr: %s", err)
	}

	return nil
}
