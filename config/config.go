package config

import (
	"github.com/gocql/gocql"
	"log"
	"os"
)

var (
	Session    *gocql.Session
	DBUser     string
	DBPassword string
	DBKeyspace string
	DBHost     string
	DBPort     string
)

func LoadConfig() {
	DBUser = os.Getenv("DB_USER")
	DBPassword = os.Getenv("DB_PASSWORD")
	DBKeyspace = os.Getenv("DB_KEYSPACE")
	DBHost = os.Getenv("DB_HOST")
	DBPort = os.Getenv("DB_PORT")

	if DBUser == "" || DBPassword == "" || DBKeyspace == "" || DBHost == "" || DBPort == "" {
		log.Fatal("Database configuration not set")
	}

	// Cassandra Cluster Setup
	cluster := gocql.NewCluster(DBHost)
	cluster.Keyspace = DBKeyspace
	cluster.Consistency = gocql.Quorum
	cluster.Port = 9042

	// Create the session
	var err error
	Session, err = cluster.CreateSession()
	if err != nil {
		log.Fatal("Failed to connect to Cassandra: ", err)
	}
	log.Println("Cassandra database connected")
}

// CloseSession Function to close the session when the app stops
func CloseSession() {
	if Session != nil {
		Session.Close()
		log.Println("Cassandra session closed")
	}
}
