package database

import (
	"encoding/json"
	"fmt"

	"github.com/cosmos/cosmos-sdk/x/group"
	"github.com/gogo/protobuf/proto"
	//dbtypes "github.com/forbole/callisto/v4/database/types"
	//"github.com/forbole/callisto/v4/types"
)

// StoreGroups stores a list of groups into the database
func (db *Db) StoreGroups(groups []*group.GroupInfo) error {
	if len(groups) == 0 {
		return nil
	}
	for _, group := range groups {
		err := db.storeGroup(group)
		if err != nil {
			return fmt.Errorf("error while storing group: %s", err)
		}
	}
	return nil
}

// StoreGroup stores a single group into the database
func (db *Db) storeGroup(group *group.GroupInfo) error {
	stmt := `
		INSERT INTO groups (id, admin_address, metadata, version, created_at)
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (id) DO UPDATE
		SET admin_address = excluded.admin_address,
			metadata = excluded.metadata,
			version = excluded.version,
			created_at = excluded.created_at`

	// Convert metadata to JSONB if it's not empty
	var metadataJSON json.RawMessage
	if group.Metadata != "" {
		metadataJSON = json.RawMessage(fmt.Sprintf(`"%s"`, group.Metadata))
	}

	_, err := db.SQL.Exec(stmt,
		group.Id,
		group.Admin,
		metadataJSON,
		group.Version,
		group.CreatedAt,
	)
	if err != nil {
		return fmt.Errorf("error while storing group: %s", err)
	}
	return nil
}

// StoreMember stores a list of group members into the database
func (db *Db) StoreGroupMembers(groupmembers []*group.GroupMember) error {
	if len(groupmembers) == 0 {
		return nil
	}
	for _, gm := range groupmembers {
		err := db.storeGroupMember(gm)
		if err != nil {
			return fmt.Errorf("error while storing group member: %s", err)
		}
	}
	return nil
}

// StoreGroup stores a single group into the database
func (db *Db) storeGroupMember(gm *group.GroupMember) error {
	stmt := `
		INSERT INTO group_members (group_id, member_address, weight, metadata, added_at)
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (group_id, member_address) DO UPDATE
		SET weight = excluded.weight,
			metadata = excluded.metadata,
			added_at = excluded.added_at`

	// Convert metadata to JSONB if it's not empty
	var metadataJSON json.RawMessage
	if gm.Member.Metadata != "" {
		metadataJSON = json.RawMessage(fmt.Sprintf(`"%s"`, gm.Member.Metadata))
	}

	_, err := db.SQL.Exec(stmt,
		gm.GroupId,
		gm.Member.Address,
		gm.Member.Weight,
		metadataJSON,
		gm.Member.AddedAt,
	)
	if err != nil {
		return fmt.Errorf("error while storing group member: %d, %s, %s", gm.GroupId, gm.Member.Address, err)
	}
	return nil
}

// StorePolicy stores a list of group policies into the database
func (db *Db) StoreGroupPolicies(gps []*group.GroupPolicyInfo) error {
	if len(gps) == 0 {
		return nil
	}
	for _, p := range gps {
		err := db.storeGroupPolicy(p)
		if err != nil {
			return fmt.Errorf("error while storing group policy: %s", err)
		}
	}
	return nil
}

// StoreGroup stores a single group into the database
func (db *Db) storeGroupPolicy(gp *group.GroupPolicyInfo) error {
	stmt := `
		INSERT INTO group_policies (address, group_id, admin_address, metadata, version, decision_policy, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		ON CONFLICT (address) DO UPDATE
		SET admin_address= excluded.admin_address,
			metadata = excluded.metadata,
			version = excluded.version,
      decision_policy = excluded.decision_policy`

	// Convert metadata to JSONB if it's not empty
	var metadataJSON json.RawMessage
	if gp.Metadata != "" {
		metadataJSON = json.RawMessage(fmt.Sprintf(`"%s"`, gp.Metadata))
	}

	// Handle decision policy
	decisionPolicyMap := make(map[string]interface{})

	// First, add the @type field
	decisionPolicyMap["@type"] = gp.DecisionPolicy.TypeUrl
	// Check the TypeUrl to determine which type to use
	switch gp.DecisionPolicy.TypeUrl {
	case "/cosmos.group.v1.ThresholdDecisionPolicy":
		var threshold group.ThresholdDecisionPolicy
		if err := proto.Unmarshal(gp.DecisionPolicy.Value, &threshold); err != nil {
			return fmt.Errorf("failed to unmarshal threshold policy: %s", err)
		}
		decisionPolicyMap["threshold"] = threshold.Threshold
		decisionPolicyMap["windows"] = map[string]string{
			"voting_period":        threshold.Windows.VotingPeriod.String(),
			"min_execution_period": threshold.Windows.MinExecutionPeriod.String(),
		}

	case "/cosmos.group.v1.PercentageDecisionPolicy":
		var percentage group.PercentageDecisionPolicy
		if err := proto.Unmarshal(gp.DecisionPolicy.Value, &percentage); err != nil {
			return fmt.Errorf("failed to unmarshal percentage policy: %s", err)
		}
		decisionPolicyMap["percentage"] = percentage.Percentage
		decisionPolicyMap["windows"] = map[string]string{
			"voting_period":        percentage.Windows.VotingPeriod.String(),
			"min_execution_period": percentage.Windows.MinExecutionPeriod.String(),
		}

	default:
		return fmt.Errorf("unknown decision policy type: %s", gp.DecisionPolicy.TypeUrl)
	}

	decisionJSON, err := json.Marshal(decisionPolicyMap)
	if err != nil {
		return fmt.Errorf("error marshalling decision policy: %s", err)
	}

	_, err = db.SQL.Exec(stmt, gp.Address, gp.GroupId, gp.Admin, metadataJSON, gp.Version, decisionJSON, gp.CreatedAt)
	if err != nil {
		return fmt.Errorf("group policy:  %s, %s", gp.Address, err)
	}
	return nil
}
