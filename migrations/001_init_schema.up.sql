CREATE EXTENSION IF NOT EXISTS ltree;


CREATE TABLE IF NOT EXISTS shelves (
    id SERIAL PRIMARY KEY,
    shelf_name VARCHAR(10),
    path ltree UNIQUE
);


CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    product_name VARCHAR(255),
    type VARCHAR(255),
    location ltree
);