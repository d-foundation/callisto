package database

import (
	"encoding/json"
	"fmt"

	dgovtypes "github.com/d-foundation/protocol/x/dgov/types"

	dbtypes "github.com/forbole/callisto/v4/database/types"
	"github.com/forbole/callisto/v4/types"
)

// SaveDGovOversightCommitteeAddress
func (db *Db) SaveDGovOCAddress(ocAddress types.DGovOversightCommitteeAddress) error {
	// SQL statement with upsert logic
	stmt := `
INSERT INTO dgov_oversight_committee_address (address, height)
VALUES ($1, $2)
ON CONFLICT (one_row_id) DO UPDATE
    SET address= excluded.address,
        height = excluded.height
WHERE dgov_oversight_committee_address.height <= excluded.height`

	// Execute the statement
	_, err := db.SQL.Exec(stmt, ocAddress.Address, ocAddress.Height)
	if err != nil {
		return fmt.Errorf("error while storing oversight committee address: %s", err)
	}
	return nil

}

// GetDGovOCAddress returns the *types.DGovOversightCommitteeAddress instance containing the current params
func (db *Db) GetDGovOCAddress() (*types.DGovOversightCommitteeAddress, error) {
	var rows []dbtypes.DGovOversightCommitteeAddressRow
	stmt := `SELECT * FROM dgov_oversight_committee_address LIMIT 1`
	err := db.Sqlx.Select(&rows, stmt)
	if err != nil {
		return nil, err
	}

	if len(rows) == 0 {
		return nil, fmt.Errorf("no staking params found")
	}

	return &types.DGovOversightCommitteeAddress{
		Address: rows[0].Address,
		Height:  rows[0].Height,
	}, nil
}

// SaveDGovOversightCommitteeParams to store the given params into the database
func (db *Db) SaveDGovOCParams(params types.DGovOversightCommitteeParams) error {
	// Marshal the params into JSON
	paramsBz, err := json.Marshal(&params.OversightCommitteeParams)
	if err != nil {
		return fmt.Errorf("error while marshaling oversight committee params: %s", err)
	}

	// SQL statement with upsert logic
	stmt := `
INSERT INTO dgov_oversight_committee_params (params, height)
VALUES ($1, $2)
ON CONFLICT (one_row_id) DO UPDATE
    SET params = excluded.params,
        height = excluded.height
WHERE dgov_oversight_committee_params.height <= excluded.height`

	// Execute the statement
	_, err = db.SQL.Exec(stmt, string(paramsBz), params.Height)
	if err != nil {
		return fmt.Errorf("error while storing oversight committee params: %s", err)
	}

	return nil
}

// GetDGovOCParams returns the types.StakingParams instance containing the current params
func (db *Db) GetDGovOCParams() (*types.DGovOversightCommitteeParams, error) {
	var rows []dbtypes.DGovOversightCommitteeParamsRow
	stmt := `SELECT * FROM dgov_oversight_committee_params LIMIT 1`
	err := db.Sqlx.Select(&rows, stmt)
	if err != nil {
		return nil, err
	}

	if len(rows) == 0 {
		return nil, fmt.Errorf("no dgov params found")
	}

	var ocParams dgovtypes.OversightCommitteeParams
	err = json.Unmarshal([]byte(rows[0].Params), &ocParams)
	if err != nil {
		return nil, err
	}

	return &types.DGovOversightCommitteeParams{
		OversightCommitteeParams: ocParams,
		Height:                   rows[0].Height,
	}, nil
}
