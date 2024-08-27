CREATE DATABASE order;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE orders (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    product_id VARCHAR(255) NOT NULL,
    quantity INT NOT NULL,
    user_id VARCHAR(255) NOT NULL
);

