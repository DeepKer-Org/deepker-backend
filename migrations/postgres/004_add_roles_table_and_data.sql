-- Create roles table
CREATE TABLE IF NOT EXISTS roles (
                                     role_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                                     role_name VARCHAR(50) UNIQUE NOT NULL
);

-- Insert initial data into roles table
INSERT INTO roles (role_name) VALUES
                                  ('admin'),
                                  ('doctor'),
                                  ('test');
