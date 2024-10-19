-- 002_add_initial_data_down.sql

-- Deleting data from 'doctor_alerts' table (Child table)
DELETE FROM public.doctor_alerts
WHERE doctor_alert_id IN (
                          'acc8f091-f376-4fc4-936e-9953cb20554d',
                          '9bf73483-db6f-4f8a-bbba-d56be1d5ba37',
                          '04dfaa2a-72c6-4fdd-945b-360f1a42e4bd',
                          '6eba1e81-aefb-43cc-94ab-f35624a5b845',
                          '78482fbb-0894-4a03-b827-464f0ebe7b99'
    );

-- Deleting data from 'doctor_patients' table (Child table)
DELETE FROM public.doctor_patients
WHERE doctor_patient_id IN (
                            'fdbadd60-9d1f-4bf6-9c16-a0ac5d3282da',
                            '2a7f893a-3269-40ad-97b1-7e8be66d1a55',
                            'cb63633a-0174-4667-8e0f-eafa7deaf6c8',
                            '44d7ef74-bf07-40f6-b3a7-69a690ce37e4',
                            'bf407595-426b-43f1-98fb-da2da22baa73',
                            '9fb7db16-719a-4c96-a250-da5a86b60032'
    );

-- Deleting data from 'computer_diagnostics' table (Child table)
DELETE FROM public.computer_diagnostics
WHERE diagnostic_id IN (
                        'fe18e71a-fb41-4a3a-85f6-0581eaaf4839',
                        'ff2f442f-0481-4d33-8702-9dd3d87d251b',
                        'c5a9b459-2283-4b53-b565-9da36eea304a'
    );

-- Deleting data from 'comorbidities' table (Child table)
DELETE FROM public.comorbidities
WHERE comorbidity_id IN (
                         '25ea8d63-7869-4ee9-9054-cf7f15b61db4',
                         'c8154392-8ddc-41cd-852a-e6ee572f1772',
                         '3b2bda86-a71d-4fbd-924d-623c97c849bd',
                         '073b9f6a-8dfb-4a17-ae1d-1613e05db40d'
    );

-- Deleting data from 'medications' table (Child table)
DELETE FROM public.medications
WHERE medication_id IN (
                        'de44255a-f7ab-425e-8a98-a24fb05a54e1',
                        'eb6da478-2708-4b85-884f-d8bfe5231dcd',
                        '47666597-c6bf-4651-9429-36405e0410fd',
                        '4f1a69f1-60d7-4ba9-b415-da4e465589bb'
    );

-- Deleting data from 'monitoring_devices' table (Parent table)
DELETE FROM public.monitoring_devices
WHERE device_id IN (
                    '13a60054-cc19-4ab4-bc32-ed833be31fa5',
                    '98fdd982-d90c-440e-818e-fdbf249a3fd1',
                    '466d20a9-6d98-45fc-a76a-80afabccb481'
    );

-- Deleting data from 'alerts' table (Parent table)
DELETE FROM public.alerts
WHERE alert_id IN (
                   '183824dd-dbb6-479e-a9cf-3f1fcc0368d9',
                   '9153da69-9fda-46a3-bdcb-fc164e47a84a',
                   '73953e79-103c-404a-81e7-83b613c648aa',
                   '50916d1c-7ef7-4790-88a5-78d4dc69f761',
                   'c6fa4d7d-d975-495e-9d46-4eb7e16485a8',
                   '3be1016c-5d6c-40f9-939a-909f47da8809',
                   '8e64f9de-9995-499a-a614-9ca8e2ea3449',
                   '864a869f-b667-4ecb-b84b-0c3e94194c3a',
                   '9639150f-7173-4fa8-bab7-d0f8a8468b94',
                   '75dd34b4-992e-4fb1-b8be-0025ae2e3893',
                   'bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb',
                   'dddddddd-dddd-dddd-dddd-dddddddddddd'
    );

-- Deleting data from 'biometric_records' table (Parent table)
DELETE FROM public.biometric_records
WHERE biometric_data_id IN (
                            '55555555-5555-5555-5555-555555555555',
                            '66666666-6666-6666-6666-666666666666',
                            'cccccccc-cccc-cccc-cccc-cccccccccccc',
                            '10df44c8-9277-4b17-bb43-0fee7094ee88'
    );

-- Deleting data from 'doctors' table (Parent table)
DELETE FROM public.doctors
WHERE doctor_id IN (
                    '44556677-8888-9999-aaaa-bbbbccccdddd',
                    '55667788-9999-aaaa-bbbb-ccccdddd1111',
                    '66778899-aaaa-bbbb-cccc-ddddeeeeffff'
    );

-- Deleting data from 'patients' table (Parent table)
DELETE FROM public.patients
WHERE patient_id IN (
                     '11111111-1111-1111-1111-111111111111',
                     '22222222-2222-2222-2222-222222222222',
                     '33333333-3333-3333-3333-333333333333',
                     'cf9ac065-293e-48e0-8798-10aab22a7282',
                     '49e6bf36-8064-4b76-9ceb-9bde332b7a57',
                     '79f76313-759b-4423-a0cf-5a23dc8c0b75',
                     '17c32d17-6d9b-4e12-aa1b-34984ae06f1c',
                     '0f25d62e-85f4-422d-bc29-a1716e8f67c0'
    );

-- Rollback (Delete records by medical_visit_id)
DELETE FROM public.medical_visits WHERE medical_visit_id IN (
                                                             '00000000-1111-1111-1111-000000000001',
                                                             '00000000-1111-1111-1111-000000000002',
                                                             '00000000-1111-1111-1111-000000000003',
                                                             '00000000-1111-1111-1111-000000000004',
                                                             '00000000-1111-1111-1111-000000000005',
                                                             '00000000-1111-1111-1111-000000000006',
                                                             '00000000-1111-1111-1111-000000000007',
                                                             '00000000-1111-1111-1111-000000000008',
                                                             '00000000-1111-1111-1111-000000000009',
                                                             '00000000-1111-1111-1111-000000000010'
    );
