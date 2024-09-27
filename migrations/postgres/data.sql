-- Insert into patients table
INSERT INTO patients (patient_id, dni, name, age, weight, height, sex, location, current_state, final_diagnosis, last_alert_id)
VALUES
    ('11111111-1111-1111-1111-111111111111', '1234567890', 'John Doe', 45, 70.5, 175.3, 'M', 'Room 101', 'Stable', 'Hypertension', NULL),
    ('22222222-2222-2222-2222-222222222222', '0987654321', 'Jane Smith', 52, 60.2, 162.4, 'F', 'Room 202', 'Critical', 'Heart Failure', NULL),
    ('33333333-3333-3333-3333-333333333333', '1122334455', 'Emily White', 30, 65.7, 168.9, 'F', 'Room 303', 'Recovering', 'Arrhythmia', NULL);

-- Insert into alerts table
INSERT INTO alerts (alert_id, alert_status, room, alert_timestamp, patient_id)
VALUES
    ('75dd34b4-992e-4fb1-b8be-0025ae2e3893', 'Unattended', 'Room 101', '2023-09-25 14:00:00', '11111111-1111-1111-1111-111111111111'),
    ('9153da69-9fda-46a3-bdcb-fc164e47a84a', 'Attended', 'Room 202', '2023-09-25 15:30:00', '22222222-2222-2222-2222-222222222222'),
    ('183824dd-dbb6-479e-a9cf-3f1fcc0368d9', 'Unattended', 'Room 303', '2023-09-26 09:15:00', '33333333-3333-3333-3333-333333333333');

-- Insert into biometrics table
INSERT INTO biometrics (biometric_id, alert_id, o2_saturation, heart_rate, systolic_blood_pressure, diastolic_blood_pressure)
VALUES
    (gen_random_uuid(), '75dd34b4-992e-4fb1-b8be-0025ae2e3893', 98, 75, 120, 80),
    (gen_random_uuid(), '9153da69-9fda-46a3-bdcb-fc164e47a84a', 95, 85, 130, 90),
    (gen_random_uuid(), '183824dd-dbb6-479e-a9cf-3f1fcc0368d9', 90, 95, 140, 100);

-- Insert into computer_diagnostics table
INSERT INTO computer_diagnostics (diagnostic_id, alert_id, diagnosis, percentage)
VALUES
    (gen_random_uuid(), '75dd34b4-992e-4fb1-b8be-0025ae2e3893', 'Possible Heart Attack', 85.50),
    (gen_random_uuid(), '9153da69-9fda-46a3-bdcb-fc164e47a84a', 'Hypertension Crisis', 90.20),
    (gen_random_uuid(), '183824dd-dbb6-479e-a9cf-3f1fcc0368d9', 'Arrhythmia', 75.80);

-- Insert into monitoring_devices table
INSERT INTO monitoring_devices (device_id, type, status, patient_id, sensors)
VALUES
    (gen_random_uuid(), 'Heart Monitor', 'In Use', '11111111-1111-1111-1111-111111111111', ARRAY['Heart Rate', 'Oxygen Level']),
    (gen_random_uuid(), 'Blood Pressure Monitor', 'Free', NULL, ARRAY['Systolic', 'Diastolic']),
    (gen_random_uuid(), 'Pulse Oximeter', 'Unavailable', '22222222-2222-2222-2222-222222222222', ARRAY['Oxygen Level']);

-- Insert into medications table
INSERT INTO medications (medication_id, patient_id, name, start_date, end_date, dosage, periodicity)
VALUES
    (gen_random_uuid(), '11111111-1111-1111-1111-111111111111', 'Aspirin', '2023-09-01', '2023-09-30', '500mg', 'Once daily'),
    (gen_random_uuid(), '22222222-2222-2222-2222-222222222222', 'Lisinopril', '2023-09-15', '2023-10-15', '20mg', 'Twice daily'),
    (gen_random_uuid(), '33333333-3333-3333-3333-333333333333', 'Metformin', '2023-09-20', '2023-10-20', '500mg', 'Once daily');

-- Insert into comorbidities table
INSERT INTO comorbidities (comorbidity_id, patient_id, comorbidity)
VALUES
    (gen_random_uuid(), '11111111-1111-1111-1111-111111111111', 'Diabetes'),
    (gen_random_uuid(), '22222222-2222-2222-2222-222222222222', 'Obesity'),
    (gen_random_uuid(), '33333333-3333-3333-3333-333333333333', 'Asthma');

-- Insert into doctors table
INSERT INTO doctors (doctor_id, dni, name, password, specialization)
VALUES
    ('44556677-8888-9999-aaaa-bbbbccccdddd', '4455667788', 'Dr. Alice Brown', 'password123', 'Cardiologist'),
    ('55667788-9999-aaaa-bbbb-ccccdddd1111', '5566778899', 'Dr. Bob Green', 'password456', 'Pulmonologist'),
    ('66778899-aaaa-bbbb-cccc-ddddeeeeffff', '6677889900', 'Dr. Charlie Blue', 'password789', 'General Practitioner');

-- Insert into doctor_alerts table
INSERT INTO doctor_alerts (doctor_alert_id, doctor_id, alert_id, attended_at)
VALUES
    (gen_random_uuid(), '44556677-8888-9999-aaaa-bbbbccccdddd', '75dd34b4-992e-4fb1-b8be-0025ae2e3893', '2023-09-25 14:30:00'),
    (gen_random_uuid(), '55667788-9999-aaaa-bbbb-ccccdddd1111', '9153da69-9fda-46a3-bdcb-fc164e47a84a', '2023-09-25 16:00:00'),
    (gen_random_uuid(), '66778899-aaaa-bbbb-cccc-ddddeeeeffff', '183824dd-dbb6-479e-a9cf-3f1fcc0368d9', '2023-09-26 09:45:00');

-- Insert into doctor_patients table
INSERT INTO doctor_patients (doctor_patient_id, doctor_id, patient_id, assigned_at)
VALUES
    (gen_random_uuid(), '44556677-8888-9999-aaaa-bbbbccccdddd', '11111111-1111-1111-1111-111111111111', '2023-09-20 10:00:00'),
    (gen_random_uuid(), '55667788-9999-aaaa-bbbb-ccccdddd1111', '22222222-2222-2222-2222-222222222222', '2023-09-21 11:00:00'),
    (gen_random_uuid(), '66778899-aaaa-bbbb-cccc-ddddeeeeffff', '33333333-3333-3333-3333-333333333333', '2023-09-22 12:00:00');
