package models

import (
	"controlF_back/internal/database"
	userModule "controlF_back/internal/modules/user"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

var DB = database.DB

func ConnectDataBase() {
	database.Connect()
	DB = database.DB

	if value, ok := os.LookupEnv("AUTO_MIGRATE"); ok && value == "true" {
		migrate()
	}
}

func migrate() {
	logrus.Info("🚀 Starting database migration...")
	err := DB.AutoMigrate(
		&Company{},
		&userModule.User{},
		&PaymentMethod{},
		&Category{},
		&Transaction{},
		&Installment{},
		&Summary{},
	)

	if err != nil {
		logrus.Fatalf("Failed to migrate database: %v", err)
	}

	logrus.Info("✅ All tables migrated successfully!")
}

func GetUserId(c *gin.Context) (uuid.UUID, error) {
	return uuid.Parse(c.GetString("x-user-id"))
}
