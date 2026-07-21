CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE admins (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- Default admin: admin@bbs.com / admin123
INSERT INTO admins (email, password_hash, name) VALUES (
    'admin@bbs.com',
    '$2a$10$v10Q0CF9hd1w2SY3VrmSQuoYY/QQzHFqEjq6rpp5zU8hZMTLhBsbW',
    'Admin BBS'
);
