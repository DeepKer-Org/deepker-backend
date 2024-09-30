-- Modificar la columna 'attended_by_id' para permitir valores nulos
ALTER TABLE alerts
    ALTER COLUMN attended_by_id DROP NOT NULL;

-- Asegurarse de que la clave foránea sigue vigente y permite valores NULL
ALTER TABLE alerts
    DROP CONSTRAINT IF EXISTS fk_attended_by,
    ADD CONSTRAINT fk_attended_by FOREIGN KEY (attended_by_id)
    REFERENCES doctors(doctor_id) ON DELETE SET NULL;
