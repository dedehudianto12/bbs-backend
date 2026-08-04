-- Renaming a product's slug used to fail because the FK below had no ON UPDATE
-- action, so the linked industry_products rows blocked the change.
ALTER TABLE industry_products
    DROP CONSTRAINT industry_products_product_slug_fkey;

ALTER TABLE industry_products
    ADD CONSTRAINT industry_products_product_slug_fkey
    FOREIGN KEY (product_slug) REFERENCES products(slug)
    ON UPDATE CASCADE ON DELETE CASCADE;
