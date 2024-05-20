CREATE TABLE IF NOT EXISTS companies (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  company_name VARCHAR(50) NOT NULL,
  company_email VARCHAR(50) UNIQUE NOT NULL,
  created_at TIMESTAMP DEFAULT current_timestamp,
  updated_at TIMESTAMP DEFAULT current_timestamp,
  deleted_at TIMESTAMP NULL
);