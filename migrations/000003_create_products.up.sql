CREATE TABLE products (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    slug VARCHAR(255) UNIQUE NOT NULL,
    name VARCHAR(255) NOT NULL,
    "group" VARCHAR(50) NOT NULL,
    kategori VARCHAR(100) NOT NULL REFERENCES categories(slug),
    category VARCHAR(255) NOT NULL,
    description TEXT NOT NULL DEFAULT '',
    detail TEXT NOT NULL DEFAULT '',
    image VARCHAR(500),
    specs JSONB NOT NULL DEFAULT '{}',
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_products_group ON products("group");
CREATE INDEX idx_products_kategori ON products(kategori);
