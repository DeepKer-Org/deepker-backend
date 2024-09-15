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
```sql
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

