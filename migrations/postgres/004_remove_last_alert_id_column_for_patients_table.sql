-- Migration: Remove last_alert_id column from patients table
ALTER TABLE patients
    DROP COLUMN last_alert_id;
