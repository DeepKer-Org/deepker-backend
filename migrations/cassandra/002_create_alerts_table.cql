CREATE TABLE IF NOT EXISTS alerts (
    alert_id UUID,
    patient_id UUID,
    room TEXT,
    alert_timestamp TIMESTAMP,
    o2_saturation INT,
    heart_rate INT,
    blood_pressure MAP<TEXT, INT>,
    computer_diagnoses LIST<TEXT>,
    alert_status TEXT,
    attended_by TEXT,
    attended_timestamp TIMESTAMP,
    final_diagnosis TEXT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    PRIMARY KEY (alert_id, patient_id)
);