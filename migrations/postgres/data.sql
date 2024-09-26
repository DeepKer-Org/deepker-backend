INSERT INTO patients (dni, name, age, weight, height, sex, location, current_state, final_diagnosis, last_alert_id)
VALUES
    ('1234567890', 'John Doe', 45, 70.5, 175.3, 'M', 'Room 101', 'Stable', 'Hypertension', NULL),
    ('0987654321', 'Jane Smith', 52, 60.2, 162.4, 'F', 'Room 202', 'Critical', 'Heart Failure', NULL),
    ('1122334455', 'Emily White', 30, 65.7, 168.9, 'F', 'Room 303', 'Recovering', 'Arrhythmia', NULL);

INSERT INTO alerts (alert_status, room, alert_timestamp, patient_id)
VALUES
    ('Unattended', 'Room 101', '2023-09-25 14:00:00', 1),
    ('Attended', 'Room 202', '2023-09-25 15:30:00', 2),
    ('Unattended', 'Room 303', '2023-09-26 09:15:00', 3);

INSERT INTO biometrics (alert_id, o2_saturation, heart_rate, systolic_blood_pressure, diastolic_blood_pressure)
VALUES
    ('75dd34b4-992e-4fb1-b8be-0025ae2e3893', 98, 75, 120, 80),
    ('9153da69-9fda-46a3-bdcb-fc164e47a84a', 95, 85, 130, 90),
    ('183824dd-dbb6-479e-a9cf-3f1fcc0368d9', 90, 95, 140, 100);

INSERT INTO computer_diagnoses (alert_id, diagnosis, percentage)
VALUES
    ('75dd34b4-992e-4fb1-b8be-0025ae2e3893', 'Possible Heart Attack', 85.50),
    ('9153da69-9fda-46a3-bdcb-fc164e47a84a', 'Hypertension Crisis', 90.20),
    ('183824dd-dbb6-479e-a9cf-3f1fcc0368d9', 'Arrhythmia', 75.80);

INSERT INTO monitoring_devices (type, status, patient_id, sensors)
VALUES
    ('Heart Monitor', 'In Use', 1, ARRAY['Heart Rate', 'Oxygen Level']),
    ('Blood Pressure Monitor', 'Free', NULL, ARRAY['Systolic', 'Diastolic']),
    ('Pulse Oximeter', 'Unavailable', 2, ARRAY['Oxygen Level']);

INSERT INTO medications (patient_id, medication, start_date, end_date, dosage, periodicity)
VALUES
    (1, 'Aspirin', '2023-09-01', '2023-09-30', '500mg', 'Once daily'),
    (2, 'Lisinopril', '2023-09-15', '2023-10-15', '20mg', 'Twice daily'),
    (3, 'Metformin', '2023-09-20', '2023-10-20', '500mg', 'Once daily');

INSERT INTO comorbidities (patient_id, comorbidity)
VALUES
    (1, 'Diabetes'),
    (2, 'Obesity'),
    (3, 'Asthma');

INSERT INTO doctors (dni, name, password, specialization)
VALUES
    ('4455667788', 'Dr. Alice Brown', 'password123', 'Cardiologist'),
    ('5566778899', 'Dr. Bob Green', 'password456', 'Pulmonologist'),
    ('6677889900', 'Dr. Charlie Blue', 'password789', 'General Practitioner');

INSERT INTO doctor_alerts (doctor_id, alert_id, attended_at)
VALUES
    (1, '75dd34b4-992e-4fb1-b8be-0025ae2e3893', '2023-09-25 14:30:00'),
    (2, '9153da69-9fda-46a3-bdcb-fc164e47a84a', '2023-09-25 16:00:00'),
    (3, '183824dd-dbb6-479e-a9cf-3f1fcc0368d9', '2023-09-26 09:45:00');

INSERT INTO doctor_patients (doctor_id, patient_id, assigned_at)
VALUES
    (1, 1, '2023-09-20 10:00:00'),
    (2, 2, '2023-09-21 11:00:00'),
    (3, 3, '2023-09-22 12:00:00');
