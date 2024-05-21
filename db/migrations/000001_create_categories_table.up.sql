CREATE TABLE IF NOT EXISTS categories (
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    name varchar(255) NOT NULL DEFAULT '',
    company_id uuid NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);
