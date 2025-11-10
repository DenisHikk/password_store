CREATE EXTENSION IF NOT EXISTS pgcrypto;

-- Main table for users
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(255) NOT NULL UNIQUE CHECK (email ~* '^[A-Za-z0-9._%-]+@[A-Za-z0-9.-]+[.][A-Za-z]+$'),
    password_hash VARCHAR(128) NOT NULL
)