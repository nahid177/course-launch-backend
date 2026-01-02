package database

import (
	"log"
	"os"

	"backend/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=" + os.Getenv("DB_HOST") +
		" user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASS") +
		" dbname=" + os.Getenv("DB_NAME") +
		" port=" + os.Getenv("DB_PORT") +
		" sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Database connection failed:", err)
	}

	// ✅ Assign global DB
	DB = db

	// ✅ Auto migrate models
	if err := DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("❌ Auto migration failed:", err)
	}

	log.Println("✅ Database connected and migrated successfully")
}
