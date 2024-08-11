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

	dsn := "host=phishing-quest.cxyy0wwccs4r.sa-east-1.rds.amazonaws.com user=labsc password=phishingquest2024 dbname=postgres port=5432 sslmode=require TimeZone=UTC connect_timeout=10"

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"host":     "phishing-quest.cxyy0wwccs4r.sa-east-1.rds.amazonaws.com",
			"user":     "labsc",
			"database": "phishingquest2024",
		}).WithError(err).Fatal("Failed to connect to database")
	}

	logrus.Info("Database connected successfully")
	return DB
}
