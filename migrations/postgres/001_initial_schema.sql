-- Patients table: Stores patient information and their latest alert reference (not a foreign key).
CREATE TABLE patients (
                          patient_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                          dni VARCHAR(10) UNIQUE NOT NULL,
                          name VARCHAR(100) NOT NULL,
                          age INT,
                          weight DECIMAL(5,2),
                          height DECIMAL(5,2),
                          sex CHAR(1),
                          location VARCHAR(100),
                          current_state VARCHAR(50),
                          final_diagnosis VARCHAR(100),
                          last_alert_id UUID,  -- Reference to the last alert, but not a foreign key.
                          created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                          updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                          deleted_at TIMESTAMP
);

-- Alerts table: Stores alert information, including patient reference.
CREATE TABLE alerts (
                        alert_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                        alert_status VARCHAR(50) NOT NULL,
                        room VARCHAR(100),
                        alert_timestamp TIMESTAMP NOT NULL,
                        attended_timestamp TIMESTAMP,
                        patient_id UUID REFERENCES patients(patient_id),
                        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                        deleted_at TIMESTAMP
);

-- Biometrics table: Stores biometric data related to an alert.
CREATE TABLE biometrics (
                            biometric_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                            alert_id UUID REFERENCES alerts(alert_id),
                            o2_saturation INT,
                            heart_rate INT,
                            systolic_blood_pressure INT,
                            diastolic_blood_pressure INT,
                            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                            updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                            deleted_at TIMESTAMP
);

-- Computer Diagnostics table: Stores automated diagnostics related to an alert.
CREATE TABLE computer_diagnostics (
                                      diagnostic_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                                      alert_id UUID REFERENCES alerts(alert_id),
                                      diagnosis VARCHAR(100) NOT NULL,
                                      percentage DECIMAL(4,2),
                                      created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                      updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                      deleted_at TIMESTAMP
);

-- Monitoring Devices table: Stores information about monitoring devices.
CREATE TABLE monitoring_devices (
                                    device_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                                    type VARCHAR(50) NOT NULL,
                                    status VARCHAR(50) CHECK (status IN ('In Use', 'Free', 'Unavailable')),
                                    patient_id UUID REFERENCES patients(patient_id),
                                    sensors TEXT[],
                                    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                    deleted_at TIMESTAMP
);

-- Medications table: Stores the medications prescribed to patients.
CREATE TABLE medications (
                             medication_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                             patient_id UUID REFERENCES patients(patient_id),
                             name VARCHAR(100) NOT NULL,
                             start_date DATE,
                             end_date DATE,
                             dosage VARCHAR(50),
                             periodicity VARCHAR(50),
                             created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                             updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                             deleted_at TIMESTAMP
);

-- Comorbidities table: Stores the comorbidities (existing conditions) of patients.
CREATE TABLE comorbidities (
                               comorbidity_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                               patient_id UUID REFERENCES patients(patient_id),
                               comorbidity VARCHAR(100) NOT NULL,
                               created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                               updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                               deleted_at TIMESTAMP
);

-- Doctors table: Stores doctor information, who log in and handle alerts and patients.
CREATE TABLE doctors (
                         doctor_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                         dni VARCHAR(10) UNIQUE NOT NULL,
                         name VARCHAR(100) NOT NULL,
                         password VARCHAR(100) NOT NULL,
                         specialization VARCHAR(100),
                         created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                         updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                         deleted_at TIMESTAMP
);

-- Doctor_Alerts table: Stores the relationship between doctors and alerts they have attended.
CREATE TABLE doctor_alerts (
                               doctor_alert_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                               doctor_id UUID REFERENCES doctors(doctor_id),
                               alert_id UUID REFERENCES alerts(alert_id),
                               attended_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Doctor_Patients table: Stores the relationship between doctors and patients they are assigned to.
CREATE TABLE doctor_patients (
                                 doctor_patient_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                                 doctor_id UUID REFERENCES doctors(doctor_id),
                                 patient_id UUID REFERENCES patients(patient_id),
                                 assigned_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
