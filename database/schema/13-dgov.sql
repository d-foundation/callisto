/* ---- PARAMS ---- */
CREATE TABLE dgov_oversight_committee_params
(
    one_row_id BOOLEAN NOT NULL DEFAULT TRUE PRIMARY KEY,
    params     JSONB   NOT NULL,
    height     BIGINT  NOT NULL,
    CHECK (one_row_id)
);

CREATE INDEX dgov_oversight_committee_params_height_index ON dgov_oversight_committee_params (height);

/* ---- COMMITTEE ---- */
CREATE TABLE dgov_oversight_committee_address
(
    one_row_id BOOLEAN NOT NULL DEFAULT TRUE PRIMARY KEY,
    address TEXT NOT NULL,
    height  BIGINT NOT NULL,
    CHECK (one_row_id)
);
CREATE INDEX dgov_oversight_committee_address_height_index ON dgov_oversight_committee_address (height);
