package main

import (
	"log"
	"os"

	"fiber-postgres-api/configs"
	"fiber-postgres-api/modules/servers"
	"fiber-postgres-api/pkg/databases"

	"github.com/joho/godotenv"
)

func main() {
	// Load dotenv config
	if err := godotenv.Load("../.env"); err != nil {
		panic(err.Error())
	}
	cfg := new(configs.Configs)

	// Fiber configs
	cfg.App.Host = os.Getenv("FIBER_HOST")
	cfg.App.Port = os.Getenv("FIBER_PORT")

	// Database Configs
	cfg.PostgreSQL.Host = os.Getenv("DB_HOST")
	cfg.PostgreSQL.Port = os.Getenv("DB_PORT")
	cfg.PostgreSQL.Protocol = os.Getenv("DB_PROTOCOL")
	cfg.PostgreSQL.Username = os.Getenv("DB_USERNAME")
	cfg.PostgreSQL.Password = os.Getenv("DB_PASSWORD")
	cfg.PostgreSQL.Database = os.Getenv("DB_DATABASE")

	// New Database
	db, err := databases.NewPostgreSQLDBConnection(cfg)
	if err != nil {
		log.Fatalln("Failed to connect to the database:", err.Error())
	}

	// Check if the connection is established
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalln("Error getting underlying database connection:", err.Error())
	}
	defer sqlDB.Close()

	if err := sqlDB.Ping(); err != nil {
		log.Fatalln("Database ping failed:", err.Error())
	} else {
		log.Println("Successfully connected to the database")
	}

	// Set up Fiber server
	s := servers.NewServer(cfg, db)
	s.Start()
}
