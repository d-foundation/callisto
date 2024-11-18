-- Groups table
CREATE TABLE groups (
    id BIGINT PRIMARY KEY,
    admin_address TEXT NOT NULL,
    metadata JSONB,
    version BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL
);

-- Group members table
CREATE TABLE group_members (
    group_id BIGINT NOT NULL,
    member_address TEXT NOT NULL,
    weight NUMERIC NOT NULL,
    metadata JSONB,
    added_at TIMESTAMP NOT NULL,
    PRIMARY KEY (group_id, member_address),
    FOREIGN KEY (group_id) REFERENCES groups(id) ON DELETE CASCADE,
    CHECK (weight >= 0)
);

-- Group policies table
CREATE TABLE group_policies (
    address TEXT PRIMARY KEY,
    group_id BIGINT NOT NULL,
    admin_address TEXT NOT NULL,
    metadata JSONB,
    version BIGINT NOT NULL,
    decision_policy JSONB NOT NULL,
    created_at TIMESTAMP NOT NULL,
    FOREIGN KEY (group_id) REFERENCES groups(id) ON DELETE CASCADE
);

-- Add indexes for common queries
CREATE INDEX idx_groups_admin ON groups(admin_address);
CREATE INDEX idx_group_policies_group_id ON group_policies(group_id);
