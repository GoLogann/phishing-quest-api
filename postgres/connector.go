package postgres

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func InitDB() *gorm.DB {

	if err := godotenv.Load(); err != nil {
		logrus.Warn("No .env file found, using default environment variables")
	}

	logrus.SetLevel(logrus.InfoLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	logrus.Info("Connecting to PostgreSQL...")

	dsn := "host=" + getEnv("DB_HOST", "phishing-quest-postgresql") +
		" user=" + getEnv("DB_USER", "default_user") +
		" password=" + getEnv("DB_PASSWORD", "default_password") +
		" dbname=" + getEnv("DB_NAME", "phishing_quest") +
		" port=" + getEnv("DB_PORT", "5432") +
		" sslmode=" + getEnv("DB_SSLMODE", "disable") +
		" TimeZone=" + getEnv("DB_TIMEZONE", "UTC") +
		" connect_timeout=" + getEnv("DB_CONNECT_TIMEOUT", "5")

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"host":     os.Getenv("DB_HOST"),
			"user":     os.Getenv("DB_USER"),
			"database": os.Getenv("DB_NAME"),
		}).WithError(err).Fatal("Failed to connect to database")
	}

	logrus.Info("Database connected successfully")
	return DB
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
