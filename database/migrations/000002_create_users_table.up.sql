CREATE TABLE IF NOT EXISTS "users"
(
    id         uuid        default gen_random_uuid() PRIMARY KEY,
    email      text                      not null,
    password   text        default null,
    created_at timestamptz default now() not null,
    updated_at timestamptz default now() not null
    -- delete_at timestamptz default null
);

-- Index
CREATE UNIQUE index users_email_index
    ON users (email);

-- Foreign Keys

-- Triggers
CREATE TRIGGER set_updated_at
    BEFORE UPDATE
    ON users
    FOR EACH ROW
EXECUTE PROCEDURE set_updated_at_to_now();

