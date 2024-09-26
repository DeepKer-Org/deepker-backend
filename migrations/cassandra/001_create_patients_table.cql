CREATE TABLE IF NOT EXISTS patients (
    patient_id UUID,
    name TEXT,
    age INT,
    current_state TEXT,
    medications LIST<TEXT>,
    created_at TIMESTAMP,
    modified_at TIMESTAMP,
    deleted_at TIMESTAMP,
    PRIMARY KEY (patient_id)
);