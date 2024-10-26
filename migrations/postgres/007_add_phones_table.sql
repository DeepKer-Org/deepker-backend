-- Create phones table
CREATE TABLE IF NOT EXISTS phones (
                                     phone_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                                     exponent_push_token VARCHAR(255) NOT NULL,
                                     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                     deleted_at TIMESTAMP
);

