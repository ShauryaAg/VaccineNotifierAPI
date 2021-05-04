package db

import (
	"fmt"
	"log"
	"os"

	"cov-api/models"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DBCon *gorm.DB
)

var (
	host     = os.Getenv("POSTGRES_HOST")
	port     = 5432
	user     = os.Getenv("POSTGRES_USER")
	password = os.Getenv("POSTGRES_PASSWORD")
	dbname   = os.Getenv("POSTGRES_DB")
)

// Initialises Database
// @returns
func CreateDatabase() (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	fmt.Println(psqlInfo)
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		log.Fatal("DB Connection failed: ", err)
	}

	migrateDatabase(db)
	return db, nil
}

func migrateDatabase(db *gorm.DB) {
	// TODO: Create migrations for models
	db.AutoMigrate(&models.User{})
}
