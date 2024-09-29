-- Add attended_by column to alerts table

ALTER TABLE patients
       ADD COLUMN attended_by UUID,
ADD CONSTRAINT fk_attended_by FOREIGN KEY (attended_by) REFERENCES doctors(doctor_id);