// Database connection and utility functions

package utils

import (
	// for database connection
	"fmt" // for getting environment variables
	"log" // for logging errors

	"github.com/jmoiron/sqlx" //for querying the database

	_ "github.com/lib/pq" // Postgres driver
)

const (
	// Database connection parameters
	DBHost     = "localhost"
	DBPort     = 5432
	DBUser     = "postgres"
	DBPassword = "yan"
	DBName     = "mydb"
)

// func ConnectDB() *sql.DB {
// 	// Get the current working directory
// 	cwd, err := os.Getwd()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Construct the database URL
// 	dbURL := os.Getenv("DATABASE_URL")
// 	if dbURL == "" {
// 		log.Fatal("DATABASE_URL environment variable is not set")
// 	}

// 	// Connect to the database
// 	db, err := sql.Open("postgres", dbURL)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return db
// }

// func ConnectDB() *sqlx.DB {
// 	// Construct the database URL
// 	dbURL := os.Getenv("DATABASE_URL")
// 	if dbURL == "" {
// 		log.Fatal("DATABASE_URL environment variable is not set")
// 	}
// 	// Connect to the database
// 	db, err := sqlx.Connect("postgres", dbURL)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return db
// }

func GetConnection() *sqlx.DB {
	// Construct the database URL
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		DBHost, DBPort, DBUser, DBPassword, DBName)

	// db, err := sqlx.Connect("postgres", psqlInfo)
	db, err := sqlx.Open("postgres", psqlInfo)

	if err != nil {
		// log.Fatal(err)
		panic(err)
	}

	log.Println("Successfully connected to the database")
	return db
}
