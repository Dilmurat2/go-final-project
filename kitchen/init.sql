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
                       created_at TIMESTAMP,
                       updated_at TIMESTAMP,
                       deleted_at TIMESTAMP,
                       is_active BOOLEAN
);