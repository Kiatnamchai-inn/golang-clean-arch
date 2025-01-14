package databases

import (
	"fiber-postgres-api/configs"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// NewPostgreSQLDBConnection creates a new PostgreSQL database connection
func NewPostgreSQLDBConnection(cfg *configs.Configs) (*gorm.DB, error) {
	// Prepare connection string
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.PostgreSQL.Host,
		cfg.PostgreSQL.Username,
		cfg.PostgreSQL.Password,
		cfg.PostgreSQL.Database,
		cfg.PostgreSQL.Port,
	)

	// Open connection to the PostgreSQL database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		// Optional GORM configurations like logging
		Logger: logger.New(
			log.New(log.Writer(), "\r\n", log.LstdFlags), // Standard log writer
			logger.Config{
				// Set to info to show all logs, change to Silent, Error, or Warn as needed
				LogLevel: logger.Info, // Levels: Silent, Error, Warn, Info
			},
		),
	})
	if err != nil {
		log.Printf("error, can't connect to database, %s", err.Error())
		return nil, err
	}

	log.Println("postgreSQL database has been connected 🐘")
	return db, nil

	// Optional: Enable auto migration (if required by your application)
	// db.AutoMigrate(&YourModelHere{})
}
