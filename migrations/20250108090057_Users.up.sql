CREATE TABLE Users (
                        id SERIAL PRIMARY KEY,
                        email VARCHAR(255) NOT NULL,
                        password VARCHAR(60) NOT NULL,
                        deleted_at TIMESTAMP DEFAULT NULL,
                        created_at TIMESTAMP NOT NULL DEFAULT NOW(),
                        updated_at TIMESTAMP NOT NULL DEFAULT NOW()
)