package postgres

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB inicializa a conexão com o banco de dados e retorna uma instância de *gorm.DB
func InitDB() *gorm.DB {
	var err error

	logrus.SetLevel(logrus.InfoLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	logrus.Info("Connecting to PostgreSQL...")

	dsn := "host=localhost user=labsc password=phishingquest dbname=phishing_quest port=5432 sslmode=disable search_path=phishing_quest"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.WithError(err).Fatal("Failed to connect to database")
	}

	logrus.Info("Database connected successfully")
	return DB
}
