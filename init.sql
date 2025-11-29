CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE if not exists users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(255) NOT NULL UNIQUE CHECK (email ~* '^[A-Za-z0-9._%-]+@[A-Za-z0-9.-]+[.][A-Za-z]+$'),
    password_hash VARCHAR(128) NOT null,
    created_at TIMESTAMPTZ NOT null default CURRENT_TIMESTAMP,
    update_at TIMESTAMPTZ NOT null default CURRENT_TIMESTAMP
)