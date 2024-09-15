INSERT INTO patients (patient_id, name, age, current_state, medications, created_at)
VALUES (uuid(), 'John Doe', 45, 'Critical', ['Aspirin', 'Metformin'], toTimestamp(now()));
INSERT INTO patients (patient_id, name, age, current_state, medications, created_at)
VALUES (uuid(), 'HDiaz', 24, 'Soft', ['Aspirin', 'Metformin'], toTimestamp(now()));
