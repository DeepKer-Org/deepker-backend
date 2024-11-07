-- Defining relationships (FOREIGN KEYS) separately
ALTER TABLE alerts
    ADD CONSTRAINT fk_patient_alert FOREIGN KEY (patient_id) REFERENCES patients(patient_id),
    ADD CONSTRAINT fk_biometric_alert FOREIGN KEY (biometric_data_id) REFERENCES biometric_data(biometric_data_id),
    ADD CONSTRAINT fk_diagnostic_alert FOREIGN KEY (diagnostic_id) REFERENCES computer_diagnostics(diagnostic_id),
    ADD CONSTRAINT fk_attended_by FOREIGN KEY (attended_by_id) REFERENCES doctors(doctor_id) ON DELETE SET NULL;

ALTER TABLE monitoring_devices
    ADD CONSTRAINT fk_patient_device FOREIGN KEY (patient_id) REFERENCES patients(patient_id),
    ADD CONSTRAINT fk_linked_by_device FOREIGN KEY (linked_by_id) REFERENCES doctors(doctor_id);

ALTER TABLE medications
    ADD CONSTRAINT fk_patient_medication FOREIGN KEY (patient_id) REFERENCES patients(patient_id);

ALTER TABLE comorbidities
    ADD CONSTRAINT fk_patient_comorbidity FOREIGN KEY (patient_id) REFERENCES patients(patient_id);

ALTER TABLE doctor_alerts
    ADD CONSTRAINT fk_doctor_alert FOREIGN KEY (doctor_id) REFERENCES doctors(doctor_id),
    ADD CONSTRAINT fk_alert_doctor_alert FOREIGN KEY (alert_id) REFERENCES alerts(alert_id);

ALTER TABLE doctor_patients
    ADD CONSTRAINT fk_doctor_patient FOREIGN KEY (doctor_id) REFERENCES doctors(doctor_id),
    ADD CONSTRAINT fk_patient_doctor_patient FOREIGN KEY (patient_id) REFERENCES patients(patient_id);

ALTER TABLE medical_visits
    ADD CONSTRAINT fk_patient_visit FOREIGN KEY (patient_id) REFERENCES patients(patient_id);