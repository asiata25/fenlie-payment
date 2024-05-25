CREATE TABLE IF NOT EXISTS detail_transactions (
  id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
  transaction_id uuid REFERENCES transactions(id),
  product_id uuid REFERENCES products(id),
  quantity int NOT NULL,
  total int NOT NULL,
  company_id uuid REFERENCES companies(id),
  created_at TIMESTAMP DEFAULT current_timestamp,
  updated_at TIMESTAMP DEFAULT current_timestamp,
  deleted_at TIMESTAMP NULL
);