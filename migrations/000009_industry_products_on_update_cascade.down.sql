ALTER TABLE industry_products
    DROP CONSTRAINT industry_products_product_slug_fkey;

ALTER TABLE industry_products
    ADD CONSTRAINT industry_products_product_slug_fkey
    FOREIGN KEY (product_slug) REFERENCES products(slug)
    ON DELETE CASCADE;
