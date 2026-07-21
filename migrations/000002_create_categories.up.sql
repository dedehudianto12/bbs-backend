CREATE TABLE categories (
    slug VARCHAR(100) PRIMARY KEY,
    label VARCHAR(255) NOT NULL,
    "group" VARCHAR(50) NOT NULL,
    sort_order INT NOT NULL DEFAULT 0
);

-- Seed 8 kategori
INSERT INTO categories (slug, label, "group", sort_order) VALUES
    ('pvc-belt',       'PVC Belt',       'belt-conveyor', 1),
    ('pu',             'PU',             'belt-conveyor', 2),
    ('flat-belt',      'Flat Belt',      'belt-conveyor', 3),
    ('rubber-belt',    'Rubber Belt',    'belt-conveyor', 4),
    ('timing-belt',    'Timing Belt',    'lainnya',       5),
    ('fastener',       'Fastener',       'lainnya',       6),
    ('cleat',          'Cleat',          'lainnya',       7),
    ('gravity-roll',   'Gravity Roll',   'lainnya',       8);
