package database

import (
	"os"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db_url := os.Getenv("DB_URL")
	var err error
	DB, err = gorm.Open(postgres.Open(db_url), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		logrus.Fatal("connection error:", err)
	} else {
		logrus.Debug("Db Connected")
	}
}
