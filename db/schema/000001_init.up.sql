-- Example initial schema for sqlc / Postgres.
-- Adjust or replace this with your real application schema.
CREATE TABLE IF NOT EXISTS guest_sessions (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP,
    expires_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS projects (
  id UUID PRIMARY KEY,
  guest_session_id UUID NULL,
  title TEXT,
  data JSONB,
  created_at TIMESTAMP
);
