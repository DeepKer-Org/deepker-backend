-- Create users table
CREATE TABLE IF NOT EXISTS users (
                                     user_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                                     email VARCHAR(100) UNIQUE NOT NULL,
                                     password VARCHAR(255) NOT NULL
);

-- Create user_roles table (relationship between users and roles)
CREATE TABLE IF NOT EXISTS user_roles (
                                          user_id UUID,
                                          role_id UUID,
                                          PRIMARY KEY (user_id, role_id),
                                          CONSTRAINT fk_user
                                              FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
                                          CONSTRAINT fk_role
                                              FOREIGN KEY (role_id) REFERENCES roles(role_id) ON DELETE CASCADE
);

-- Insert example data into users table
INSERT INTO users (email, password) VALUES
                                                  ('admin@example.com', 'hashed_password1'),
                                                  ('doctor@example.com', 'hashed_password2'),
                                                  ('user@example.com', 'hashed_password3');

-- Insert example data into user_roles table (assign roles to users)
INSERT INTO user_roles (user_id, role_id) VALUES
                                              ((SELECT user_id FROM users WHERE email = 'admin@example.com'), (SELECT role_id FROM roles WHERE role_name = 'admin')),
                                              ((SELECT user_id FROM users WHERE email = 'doctor@example.com'), (SELECT role_id FROM roles WHERE role_name = 'doctor')),
                                              ((SELECT user_id FROM users WHERE email = 'user@example.com'), (SELECT role_id FROM roles WHERE role_name = 'test')),
                                              ((SELECT user_id FROM users WHERE email = 'admin@example.com'), (SELECT role_id FROM roles WHERE role_name = 'doctor')); -- Example of user with multiple roles
