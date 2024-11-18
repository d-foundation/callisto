package types

// DGovOversightCommitteeParams a single row inside the dgov_oversight_committee table
type DGovOversightCommitteeParamsRow struct {
	OneRowID bool   `db:"one_row_id"`
	Params   string `db:"params"`
	Height   int64  `db:"height"`
}

type DGovOversightCommitteeAddressRow struct {
	OneRowID bool   `db:"one_row_id"`
	Address  string `db:"address"`
	Height   int64  `db:"height"`
}
