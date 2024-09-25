INSERT INTO patients (dni, name, age, weight, height, sex, location, current_state, final_diagnosis)
VALUES
    ('1234567890', 'John Doe', 45, 75.50, 1.75, 'M', 'Lima', 'Estable', 'Hipertensión'),
    ('0987654321', 'Jane Smith', 60, 68.00, 1.65, 'F', 'Cusco', 'Crítico', 'Insuficiencia cardíaca');

INSERT INTO alerts (alert_id, alert_status, attended_by, alert_timestamp, patient_id)
VALUES
    ('9f7b1f83-1ec6-4c5f-8af9-6d4f6c074d39' , 'Active', 'Dr. Garcia', '2024-09-23 14:32:00', 1),
    ('8e4c926d-66f7-4e82-9b11-86eeb2f63c0a', 'Resolved', 'Dr. Perez', '2024-09-22 11:20:00', 2);


INSERT INTO biometrics (alert_id, o2_saturation, heart_rate, systolic_blood_pressure, diastolic_blood_pressure)
VALUES
    ('9f7b1f83-1ec6-4c5f-8af9-6d4f6c074d39', 95, 75, 120, 80),
    ('8e4c926d-66f7-4e82-9b11-86eeb2f63c0a', 89, 95, 145, 90);

INSERT INTO computer_diagnoses (alert_id, diagnosis, percentage)
VALUES
    ('9f7b1f83-1ec6-4c5f-8af9-6d4f6c074d39', 'Hipertensión probable', 85.00),
    ('8e4c926d-66f7-4e82-9b11-86eeb2f63c0a', 'Fallo cardíaco inminente', 92.50);


INSERT INTO alert_doctors (alert_id, doctor_name)
VALUES
    ('9f7b1f83-1ec6-4c5f-8af9-6d4f6c074d39', 'Dr. Garcia'),
    ('8e4c926d-66f7-4e82-9b11-86eeb2f63c0a', 'Dr. Perez');


INSERT INTO patient_doctors (patient_id, doctor_name)
VALUES
    (1, 'Dr. Garcia'),
    (2, 'Dr. Perez');

INSERT INTO comorbidities (patient_id, comorbidity)
VALUES
    (1, 'Diabetes'),
    (2, 'Colesterol alto');

INSERT INTO medications (patient_id, medication)
VALUES
    (1, 'Metformina'),
    (2, 'Atorvastatina');
