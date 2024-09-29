ALTER TABLE alerts
    ADD COLUMN attended_by_id UUID,
    ADD CONSTRAINT fk_attended_by FOREIGN KEY (attended_by_id) REFERENCES doctors(doctor_id);
