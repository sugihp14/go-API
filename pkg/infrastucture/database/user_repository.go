// pkg/infrastructure/database/user_repository.go

package database

import (
	domain "go-crud/pkg/domain/user"

	"gorm.io/gorm"
)

type UserRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
    return &UserRepository{
        db: db,
    }
}

func (r *UserRepository) CreateUser(user *domain.User) error {
    return r.db.Create(user).Error
}

func (r *UserRepository) GetUserByUsername(username string) (*domain.User, error) {
    var user domain.User
    err := r.db.Where("username = ?", username).First(&user).Error
    if err != nil {
        return nil, err
    }
    return &user, nil
}
