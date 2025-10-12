package models

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDataBase() {
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

	if value, ok := os.LookupEnv("AUTO_MIGRATE"); ok && value == "true" {
		migrate()
	}
}

func migrate() {
	DB.AutoMigrate(
		&User{},
		&Company{},
		&Transaction{},
		&Installment{},
		&Category{},
		&Summary{},
	)
}

func GetUserId(c *gin.Context) (uuid.UUID, error) {
	return uuid.Parse(c.GetString("x-user-id"))
}
