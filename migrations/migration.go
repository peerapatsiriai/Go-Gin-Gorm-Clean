package migrations

import (
	entities "go/api/internal/entities"

	"gorm.io/gorm"
)

func MigrateRun(db *gorm.DB) error {
	return db.AutoMigrate(&entities.Book{})
}
