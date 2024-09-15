## create the container
```bash
docker run -d --name cassandra-container -p 9042:9042 cassandra:latest
```
## connect to the container
```bash
docker exec -it cassandra-container cqlsh
```
## create a keyspace
```sql
CREATE KEYSPACE deepker WITH REPLICATION = {'class' : 'SimpleStrategy', 'replication_factor' : 1};
```
## use the keyspace
```sql
USE deepker;
```
## create a table
```cassandraql
CREATE TABLE patients (
    id UUID PRIMARY KEY,
    name TEXT,
    age INT,
    current_state TEXT,
    medications LIST<TEXT>,
    created_at TIMESTAMP
);
```
## describe the table
```sql
DESCRIBE TABLE patients;
```
## insert data
```sql
INSERT INTO patients (id, name, age, current_state, medications, created_at)
VALUES (uuid(), 'John Doe', 45, 'Critical', ['Aspirin', 'Metformin'], toTimestamp(now()));
```
## select data
```sql
SELECT * FROM patients;
```

## create the table for alerts
```cassandraql
CREATE TABLE IF NOT EXISTS alerts (
    alert_id UUID,
    patient_id UUID,
    room TEXT,
    alert_timestamp TIMESTAMP,
    o2_saturation INT,
    heart_rate INT,
    blood_pressure MAP<TEXT, INT>,
    computer_diagnoses LIST<TEXT>,
    alert_status TEXT,
    attended_by TEXT,
    attended_timestamp TIMESTAMP,
    final_diagnosis TEXT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    PRIMARY KEY (alert_id, patient_id)
);
```
## insert data
```sql
INSERT INTO alerts (alert_id, patient_id, room, alert_timestamp, o2_saturation, heart_rate, blood_pressure, computer_diagnoses, alert_status, attended_by, attended_timestamp, final_diagnosis, created_at, updated_at)
VALUES (uuid(), uuid(), 'ICU', toTimestamp(now()), 95, 120, {'systolic': 120, 'diastolic': 80}, ['Hypertension'], 'Open', 'Dr. Smith', toTimestamp(now()), 'Hypertension', toTimestamp(now()), toTimestamp(now()));
```
```cassandraql
INSERT INTO alerts (alert_id, patient_id, room, alert_timestamp, o2_saturation, heart_rate, blood_pressure, computer_diagnoses, alert_status, attended_by, attended_timestamp, final_diagnosis, created_at, updated_at)
VALUES (uuid(), f23955a0-736d-11ef-87b3-2a024c401a50, 'ICU', toTimestamp(now()), 95, 120, {'systolic': 120, 'diastolic': 80}, ['Hypertension'], 'Open', 'Dr. Smith', toTimestamp(now()), 'Hypertension', toTimestamp(now()), toTimestamp(now()));
INSERT INTO alerts (alert_id, patient_id, room, alert_timestamp, o2_saturation, heart_rate, blood_pressure, computer_diagnoses, alert_status, attended_by, attended_timestamp, final_diagnosis, created_at, updated_at)
VALUES (uuid(), f23955a0-736d-11ef-87b3-2a024c401a50, 'ICU', toTimestamp(now()), 95, 120, {'systolic': 120, 'diastolic': 80}, ['Hypertension'], 'Open', 'Dr. Smith', toTimestamp(now()), 'Hypertension', toTimestamp(now()), toTimestamp(now()));

```


