package migrations

import (
	domain "go-crud/pkg/domain/user"

	"gorm.io/gorm"
)

func CreateUserTable(db *gorm.DB) error {
	if db.Migrator().HasTable(&domain.User{}) {
		// Table already exists, no need to create
		return nil
	}
	return db.AutoMigrate(&domain.User{})
}