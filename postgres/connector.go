package postgres

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB inicializa a conexão com o banco de dados e retorna uma instância de *gorm.DB
func InitDB() *gorm.DB {
	var err error
	fmt.Println("Connecting to PostgreSQL...")
	fmt.Println(err)
	dsn := "host=localhost user=user password=password dbname=phishing_quest port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	//if err != nil {
	//	log.Fatal("failed to connect database:", err)
	//}
	//log.Println("database connected")
	return DB
}
