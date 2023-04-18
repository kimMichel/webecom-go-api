package migrations

import (
	"github.com/kimMichel/webecom-go-api/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Product{})
}
