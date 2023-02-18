-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE person(
    id BIGINT NOT NULL,
    first_name VARCHAR(255),
    last_name VARCHAR(255)
)

-- +migrate StatementEnd