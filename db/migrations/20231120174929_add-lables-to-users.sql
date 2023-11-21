-- migrate:up
ALTER TABLE users ADD COLUMN labels VARCHAR[] NOT NULL DEFAULT '{}';

-- migrate:down
ALTER TABLE users DROP COLUMN labels;

