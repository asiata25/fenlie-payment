ALTER TABLE users ALTER COLUMN "role" TYPE user_type USING "role"::text::user_type;
