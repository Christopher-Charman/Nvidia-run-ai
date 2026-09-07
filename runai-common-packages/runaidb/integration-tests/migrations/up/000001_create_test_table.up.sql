CREATE TABLE IF NOT EXISTS test_secrets (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    secret TEXT NOT NULL
);

-- Schema the service user shouldn't have access to
CREATE SCHEMA IF NOT EXISTS restricted_schema;
CREATE TABLE IF NOT EXISTS restricted_schema.restricted_table (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);
