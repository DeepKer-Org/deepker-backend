-- 003_foreign_keys_down.sql

-- Dropping foreign key constraints from 'doctor_patients' table
ALTER TABLE doctor_patients
    DROP CONSTRAINT IF EXISTS fk_doctor_patient,
    DROP CONSTRAINT IF EXISTS fk_patient_doctor_patient;

-- Dropping foreign key constraints from 'doctor_alerts' table
ALTER TABLE doctor_alerts
    DROP CONSTRAINT IF EXISTS fk_doctor_alert,
    DROP CONSTRAINT IF EXISTS fk_alert_doctor_alert;

-- Dropping foreign key constraint from 'comorbidities' table
ALTER TABLE comorbidities
    DROP CONSTRAINT IF EXISTS fk_patient_comorbidity;

-- Dropping foreign key constraint from 'medications' table
ALTER TABLE medications
    DROP CONSTRAINT IF EXISTS fk_patient_medication;

-- Dropping foreign key constraint from 'monitoring_devices' table
ALTER TABLE monitoring_devices
    DROP CONSTRAINT IF EXISTS fk_patient_device,
    DROP CONSTRAINT IF EXISTS fk_linked_by_device;

-- Dropping foreign key constraint from 'computer_diagnostics' table
ALTER TABLE computer_diagnostics
    DROP CONSTRAINT IF EXISTS fk_alert_diagnostic;

-- Dropping foreign key constraints from 'alerts' table
ALTER TABLE alerts
    DROP CONSTRAINT IF EXISTS fk_patient_alert,
    DROP CONSTRAINT IF EXISTS fk_biometric_alert,
    DROP CONSTRAINT IF EXISTS fk_attended_by;