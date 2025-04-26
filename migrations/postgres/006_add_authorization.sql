ALTER TABLE "doctors"
    ADD COLUMN user_id UUID;

UPDATE "doctors"
SET user_id = (SELECT user_id FROM "users" WHERE username = '44556677')
WHERE doctor_id = '44556677-8888-9999-aaaa-bbbbccccdddd';

UPDATE "doctors"
SET user_id = (SELECT user_id FROM "users" WHERE username = '55667788')
WHERE doctor_id = '55667788-9999-aaaa-bbbb-ccccdddd1111';

UPDATE "doctors"
SET user_id = (SELECT user_id FROM "users" WHERE username = '66778899')
WHERE doctor_id = '66778899-aaaa-bbbb-cccc-ddddeeeeffff';


ALTER TABLE "doctors"
    ALTER COLUMN user_id SET NOT NULL;

ALTER TABLE "doctors"
    ADD CONSTRAINT fk_user_doctor
        FOREIGN KEY (user_id)
            REFERENCES "users" (user_id);
