package postgres

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

// InitDB inicializa a conexão com o banco de dados e retorna uma instância de *gorm.DB
func InitDB() *gorm.DB {
	var err error
	fmt.Println("Connecting to PostgreSQL...")

	dsn := "host=localhost user=labsc password=phishingquest dbname=phishing_quest port=5432 sslmode=disable search_path=phishing_quest"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	log.Println("Database connected successfully")
	return DB
}
