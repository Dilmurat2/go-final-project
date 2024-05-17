CREATE TABLE IF NOT EXISTS orders (
                        id varchar PRIMARY KEY,
                        status VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS items (
                       id SERIAL PRIMARY KEY,
                       order_id varchar REFERENCES orders(id) ON DELETE CASCADE,
                       name VARCHAR(255) NOT NULL,
                       price FLOAT NOT NULL,
                       weight FLOAT NOT NULL,
                       created_at TIMESTAMP,
                       updated_at TIMESTAMP,
                       deleted_at TIMESTAMP,
                       is_active BOOLEAN NOT NULL
);