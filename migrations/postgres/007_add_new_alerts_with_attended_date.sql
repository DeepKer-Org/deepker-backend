INSERT INTO alerts (alert_id, room, alert_timestamp, patient_id, biometric_data_id, attended_timestamp)
VALUES
    ('zzzzzzzz-zzzz-zzzz-zzzz-zzzzzzzzzzzz', 'Room 301', '2023-09-25 14:00:00', '11111111-1111-1111-1111-111111111111', '55555555-5555-5555-5555-555555555555', '2023-09-25 14:30:00'),
    ('bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb', 'Room 402', '2023-09-25 15:30:00', '22222222-2222-2222-2222-222222222222', '66666666-6666-6666-6666-666666666666', '2023-09-25 16:00:00'),;

INSERT INTO doctor_alerts (doctor_alert_id, doctor_id, alert_id, attended_at)
VALUES
    (gen_random_uuid(), '44556677-8888-9999-aaaa-bbbbccccdddd', 'zzzzzzzz-zzzz-zzzz-zzzz-zzzzzzzzzzzz', '2023-09-25 14:30:00'),
    (gen_random_uuid(), '55667788-9999-aaaa-bbbb-ccccdddd1111', 'bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb', '2023-09-25 16:00:00');