-- +goose Up
CREATE TABLE orders (
                        id varchar PRIMARY KEY,
                        status VARCHAR(255) NOT NULL
);

CREATE TABLE items (
                       id SERIAL PRIMARY KEY,
                       order_id varchar REFERENCES orders(id) ON DELETE CASCADE,
                       name VARCHAR(255),
                       price FLOAT,
                       weight FLOAT,
                       created_at varchar(255),
                       updated_at varchar(255),
                       deleted_at varchar(255),
                       is_active BOOLEAN
);

-- +goose Down
DROP TABLE items;
DROP TABLE orders;