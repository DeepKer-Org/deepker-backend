package utils

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// TODO: Compare hashes to see if any migrations were changed.
// TODO: Reuse the configuration connection
func ExecuteMigrations() {
	// Parse command-line flags
	resetFlag := flag.Bool("reset", false, "Reset the database by rolling back all migrations")
	flag.Parse()

	// Load environment variables from .env file or use defaults
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using default environment variables")
	}

	// Build the connection string
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	sslmode := os.Getenv("SSL_MODE")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbHost, dbPort, dbUser, dbPassword, dbName, sslmode)

	// Open the database connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	// Ensure the database connection is alive
	err = db.Ping()
	if err != nil {
		log.Fatalf("Cannot reach the database: %v", err)
	}

	log.Println("Connected to the database successfully!")

	// Initialize the migrations table if it doesn't exist
	err = createMigrationsTable(db)
	if err != nil {
		log.Fatalf("Error creating migrations table: %v", err)
	}

	// If reset flag is set, roll back all migrations
	if *resetFlag {
		log.Println("Reset flag detected. Rolling back all migrations...")
		err = rollbackMigrations(db)
		if err != nil {
			log.Fatalf("Error rolling back migrations: %v", err)
		}
		log.Println("All migrations rolled back successfully!")
		return
	}

	// Get the list of migration files
	migrationDir := filepath.Join("migrations", "postgres")
	files, err := ioutil.ReadDir(migrationDir)
	if err != nil {
		log.Fatalf("Error reading migrations directory: %v", err)
	}

	// Filter and sort SQL files for up migrations
	var migrationFiles []string
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".sql") && !strings.Contains(file.Name(), "_down.sql") {
			migrationFiles = append(migrationFiles, file.Name())
		}
	}
	sort.Strings(migrationFiles) // Ensure the files are executed in order

	// Get the list of applied migrations from the database
	appliedMigrations, err := getAppliedMigrations(db)
	if err != nil {
		log.Fatalf("Error fetching applied migrations: %v", err)
	}

	// Apply pending migrations
	for _, fileName := range migrationFiles {
		if _, applied := appliedMigrations[fileName]; applied {
			log.Printf("Skipping already applied migration: %s", fileName)
			continue
		}

		filePath := filepath.Join(migrationDir, fileName)
		log.Printf("Executing migration: %s", fileName)

		err := executeSQLFile(db, filePath)
		if err != nil {
			log.Fatalf("Error executing migration %s: %v", fileName, err)
		}

		// Record the applied migration
		err = recordMigration(db, fileName)
		if err != nil {
			log.Fatalf("Error recording migration %s: %v", fileName, err)
		}
	}

	log.Println("All pending migrations executed successfully!")
}

func createMigrationsTable(db *sql.DB) error {
	query := `
    CREATE TABLE IF NOT EXISTS schema_migrations (
        version VARCHAR(255) PRIMARY KEY,
        applied_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
    );`
	_, err := db.Exec(query)
	return err
}

func getAppliedMigrations(db *sql.DB) (map[string]bool, error) {
	rows, err := db.Query("SELECT version FROM schema_migrations")
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	appliedMigrations := make(map[string]bool)
	for rows.Next() {
		var version string
		err := rows.Scan(&version)
		if err != nil {
			return nil, err
		}
		appliedMigrations[version] = true
	}

	return appliedMigrations, nil
}

func recordMigration(db *sql.DB, version string) error {
	_, err := db.Exec("INSERT INTO schema_migrations (version) VALUES ($1)", version)
	return err
}

func removeMigrationRecord(db *sql.DB, version string) error {
	_, err := db.Exec("DELETE FROM schema_migrations WHERE version = $1", version)
	return err
}

func executeSQLFile(db *sql.DB, filePath string) error {
	// Read the SQL file
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error reading SQL file: %v", err)
	}

	// Split the content into individual statements
	statements := strings.Split(string(content), ";")

	// Execute each statement within a transaction
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("error beginning transaction: %v", err)
	}

	for _, stmt := range statements {
		stmt = strings.TrimSpace(stmt)
		if stmt == "" {
			continue
		}
		_, err := tx.Exec(stmt)
		if err != nil {
			err := tx.Rollback()
			if err != nil {
				return err
			}
			return fmt.Errorf("error executing statement: %v\nStatement: %s", err, stmt)
		}
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("error committing transaction: %v", err)
	}

	return nil
}

func rollbackMigrations(db *sql.DB) error {
	// Get the list of applied migrations, ordered by applied_at descending
	rows, err := db.Query("SELECT version FROM schema_migrations ORDER BY applied_at DESC")
	if err != nil {
		return err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var appliedMigrations []string
	for rows.Next() {
		var version string
		err := rows.Scan(&version)
		if err != nil {
			return err
		}
		appliedMigrations = append(appliedMigrations, version)
	}

	// Rollback migrations in reverse order
	for _, version := range appliedMigrations {
		// Find the corresponding down migration file
		downFileName := strings.Replace(version, ".sql", "_down.sql", 1)
		downFilePath := filepath.Join("migrations", "postgres", downFileName)

		// Check if the down migration file exists
		if _, err := os.Stat(downFilePath); os.IsNotExist(err) {
			log.Printf("No down migration found for %s. Skipping.", version)
			continue
		}

		log.Printf("Rolling back migration: %s", version)

		err := executeSQLFile(db, downFilePath)
		if err != nil {
			return fmt.Errorf("error executing down migration %s: %v", downFileName, err)
		}

		// Remove the migration record
		err = removeMigrationRecord(db, version)
		if err != nil {
			return fmt.Errorf("error removing migration record %s: %v", version, err)
		}
	}

	return nil
}
