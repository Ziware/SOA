CREATE SCHEMA IF NOT EXISTS users;

CREATE TABLE IF NOT EXISTS users.users (
    user_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    login VARCHAR(30) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL UNIQUE,
    password_hash VARCHAR(256) NOT NULL,
    name VARCHAR(100),
    surname VARCHAR(100),
    birth_date VARCHAR(16),
    phone_number VARCHAR(20),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_user_id ON users.users(user_id);
CREATE INDEX IF NOT EXISTS idx_users_login ON users.users(login);
CREATE INDEX IF NOT EXISTS idx_users_email ON users.users(email);

CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_update_users_updated_at
BEFORE UPDATE ON users.users
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

CREATE TABLE IF NOT EXISTS users.user_sessions (
    user_id UUID NOT NULL REFERENCES users.users(user_id) ON DELETE CASCADE,
    token VARCHAR(512) NOT NULL,
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_user_sessions_token ON users.user_sessions(token);
CREATE INDEX IF NOT EXISTS idx_user_sessions_user_id ON users.user_sessions(user_id);
