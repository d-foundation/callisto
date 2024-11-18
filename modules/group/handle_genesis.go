package group

import (
	"encoding/json"
	"fmt"

	tmtypes "github.com/cometbft/cometbft/types"

	group "github.com/cosmos/cosmos-sdk/x/group"
	"github.com/rs/zerolog/log"
)

// HandleGenesis implements modules.Module
func (m *Module) HandleGenesis(doc *tmtypes.GenesisDoc, appState map[string]json.RawMessage) error {
	log.Debug().Str("module", "group").Msg("parsing genesis")

	// Read the genesis state
	var genStateGroup group.GenesisState
	err := m.cdc.UnmarshalJSON(appState[group.ModuleName], &genStateGroup)
	if err != nil {
		return fmt.Errorf("error while reading gov genesis data: %s", err)
	}

	// Save the groups
	err = m.db.StoreGroups(genStateGroup.Groups)
	if err != nil {
		return fmt.Errorf("error while storing genesis groups: %s", err)
	}

	// Save the group members
	err = m.db.StoreGroupMembers(genStateGroup.GroupMembers)
	if err != nil {
		return fmt.Errorf("error while storing genesis group members: %s", err)
	}

	// Save the group policies
	err = m.db.StoreGroupPolicies(genStateGroup.GroupPolicies)
	if err != nil {
		return fmt.Errorf("error while storing genesis group policies: %s", err)
	}

	return nil
}
