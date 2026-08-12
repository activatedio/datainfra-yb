-- +goose Up

CREATE EXTENSION IF NOT EXISTS pg_trgm;

CREATE TABLE categories (
    name VARCHAR(64),
    description VARCHAR(200),
    PRIMARY KEY (name)
);

CREATE TABLE products (
    sku VARCHAR(64),
    description VARCHAR(200),
    full_text TSVECTOR GENERATED ALWAYS AS (TO_TSVECTOR
      ('english',
        (CASE WHEN description IS NULL THEN '' ELSE description END)
      )
    ) STORED,
    PRIMARY KEY (sku)
);

CREATE INDEX products_description_trgm ON products USING ybgin (description gin_trgm_ops);

CREATE TABLE product_categories (
    product_sku VARCHAR(64),
    category_name VARCHAR(64),
    created_at TIMESTAMP NOT NULL,
    PRIMARY KEY (product_sku, category_name),
    FOREIGN KEY (product_sku) REFERENCES products(sku),
    FOREIGN KEY (category_name) REFERENCES categories(name)
);

CREATE TABLE themes2 (
    tenant_id VARCHAR(64),
    name VARCHAR(64),
    description VARCHAR(200),
    PRIMARY KEY (tenant_id, name)
);

CREATE TABLE locations (
    city VARCHAR(64),
    state VARCHAR(64),
    latitude REAL,
    longitude REAL,
    PRIMARY KEY (city, state)
);
