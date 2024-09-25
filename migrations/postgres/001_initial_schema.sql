-- Tabla patients
CREATE TABLE patients (
                          patient_id SERIAL PRIMARY KEY,
                          dni VARCHAR(10) UNIQUE NOT NULL,
                          name VARCHAR(100) NOT NULL,
                          age INT,
                          weight DECIMAL(5,2),
                          height DECIMAL(5,2),
                          sex CHAR(1),
                          location VARCHAR(100),
                          current_state VARCHAR(50),
                          final_diagnosis VARCHAR(100),
                          created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                          updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                          deleted_at TIMESTAMP
);

-- Tabla alerts
CREATE TABLE alerts (
                        alert_id UUID PRIMARY KEY,
                        alert_status VARCHAR(50) NOT NULL,
                        attended_by VARCHAR(100),
                        alert_timestamp TIMESTAMP NOT NULL,
                        attended_timestamp TIMESTAMP,
                        patient_id INT REFERENCES patients(patient_id),
                        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                        deleted_at TIMESTAMP
);

-- Tabla biometrics
CREATE TABLE biometrics (
                            alert_id UUID PRIMARY KEY REFERENCES alerts(alert_id),
                            o2_saturation INT,
                            heart_rate INT,
                            systolic_blood_pressure INT,
                            diastolic_blood_pressure INT,
                            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                            updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                            deleted_at TIMESTAMP
);

-- Tabla computer_diagnoses
CREATE TABLE computer_diagnoses (
                                    id SERIAL PRIMARY KEY,
                                    alert_id UUID REFERENCES alerts(alert_id),
                                    diagnosis VARCHAR(100) NOT NULL,
                                    percentage DECIMAL(4,2),
                                    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                    deleted_at TIMESTAMP
);

-- Tabla alert_doctors
CREATE TABLE alert_doctors (
                               id SERIAL PRIMARY KEY,
                               alert_id UUID REFERENCES alerts(alert_id),
                               doctor_name VARCHAR(100) NOT NULL,
                               created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                               updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                               deleted_at TIMESTAMP
);

-- Tabla patient_doctors
CREATE TABLE patient_doctors (
                                 id SERIAL PRIMARY KEY,
                                 patient_id INT REFERENCES patients(patient_id),
                                 doctor_name VARCHAR(100) NOT NULL,
                                 created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                 updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                 deleted_at TIMESTAMP
);

-- Tabla comorbidities
CREATE TABLE comorbidities (
                               id SERIAL PRIMARY KEY,
                               patient_id INT REFERENCES patients(patient_id),
                               comorbidity VARCHAR(100) NOT NULL,
                               created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                               updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                               deleted_at TIMESTAMP
);

-- Tabla medications
CREATE TABLE medications (
                             id SERIAL PRIMARY KEY,
                             patient_id INT REFERENCES patients(patient_id),
                             medication VARCHAR(100) NOT NULL,
                             created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                             updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                             deleted_at TIMESTAMP
);
