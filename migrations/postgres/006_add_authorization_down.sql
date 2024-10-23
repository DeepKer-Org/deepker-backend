-- Remove the foreign key constraint
ALTER TABLE "doctors"
    DROP CONSTRAINT IF EXISTS fk_user_doctor;

-- Allow user_id to be null again
ALTER TABLE "doctors"
    ALTER COLUMN user_id DROP NOT NULL;

-- Drop the user_id column from doctors
ALTER TABLE "doctors"
    DROP COLUMN IF EXISTS user_id;
