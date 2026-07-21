CREATE TABLE industries (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    slug VARCHAR(255) UNIQUE NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL DEFAULT '',
    image VARCHAR(500),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE industry_products (
    industry_id UUID NOT NULL REFERENCES industries(id) ON DELETE CASCADE,
    product_slug VARCHAR(255) NOT NULL REFERENCES products(slug) ON DELETE CASCADE,
    PRIMARY KEY (industry_id, product_slug)
);
