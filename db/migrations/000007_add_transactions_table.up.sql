CREATE TABLE IF NOT EXISTS transactions (
  id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
  status payment_type,
  total int NOT NULL,
  company_id uuid REFERENCES companies(id) NOT NULL,
  user_id uuid REFERENCES companies(id) NOT NULL,
  order_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT current_timestamp,
  deleted_at TIMESTAMP NULL
);