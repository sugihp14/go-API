package migrations

import (
	domain "go-crud/pkg/domain/post"

	"gorm.io/gorm"
)

func CreatePostTable(db *gorm.DB) error {
	if db.Migrator().HasTable(&domain.Post{}) {
		// Table already exists, no need to create
		return nil
	}
	return db.AutoMigrate(&domain.Post{})
}