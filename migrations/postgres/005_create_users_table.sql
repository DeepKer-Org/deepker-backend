-- Create users table
CREATE TABLE IF NOT EXISTS users (
                                     user_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                                     username VARCHAR(100) UNIQUE NOT NULL,
                                     password VARCHAR(255) NOT NULL,
                                     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                     deleted_at TIMESTAMP
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
INSERT INTO users (username, password) VALUES
                                                  ('44556677', '$2a$10$tBhrdwxV2Hc1jHyRxBdgve3PL/GlIr5YTDV3O0KBIbbHRbdpGtTzS'),
                                                  ('55667788', '$2a$10$h5.2YfnHQs/CgyRMyew.j.DCiJMbVwSPlAAuWmgEbNcwNAyp4vrVe'),
                                                  ('66778899', '$2a$10$qi.n3zf5C5VBpgcrehsTW.j.7Xu5JZjRhivCha8R6PUd8cl/rU9Ni'),
                                                  ('admin@example.com', '$2a$10$tBhrdwxV2Hc1jHyRxBdgve3PL/GlIr5YTDV3O0KBIbbHRbdpGtTzS');

-- Insert example data into user_roles table (assign roles to users)
INSERT INTO user_roles (user_id, role_id) VALUES
                                              ((SELECT user_id FROM users WHERE username = '44556677'), (SELECT role_id FROM roles WHERE role_name = 'doctor')),
                                              ((SELECT user_id FROM users WHERE username = '55667788'), (SELECT role_id FROM roles WHERE role_name = 'doctor')),
                                              ((SELECT user_id FROM users WHERE username = '66778899'), (SELECT role_id FROM roles WHERE role_name = 'doctor')),
                                              ((SELECT user_id FROM users WHERE username = 'admin@example.com'), (SELECT role_id FROM roles WHERE role_name = 'admin'));
