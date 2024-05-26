CREATE TABLE IF NOT EXISTS invoices (
  id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
  email_customer varchar(50) NOT NULL,
  transaction_id uuid REFERENCES transactions(id) NOT NULL,
  payment_method varchar(50),
  amount int NOT NULL,
  status payment_type DEFAULT unpaid,
  company_id uuid REFERENCES companies(id),
  created_at TIMESTAMP DEFAULT current_timestamp,
  updated_at TIMESTAMP DEFAULT current_timestamp,
  deleted_at TIMESTAMP NULL
);