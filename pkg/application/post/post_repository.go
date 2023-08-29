package post

import (
	"go-crud/pkg/domain/post"

	"gorm.io/gorm"
)

type Repository interface {
	Create(post *post.Post) error
	Update(post *post.Post) error
	Delete(post *post.Post) error
	GetByID(id uint) (*post.Post, error)
	GetAll() ([]*post.Post, error)
}

type postRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &postRepository{
		db: db,
	}
}

func (r *postRepository) Create(post *post.Post) error {
	return r.db.Create(post).Error
}

func (r *postRepository) Update(post *post.Post) error {
	return r.db.Save(post).Error
}

func (r *postRepository) Delete(post *post.Post) error {
	return r.db.Delete(post).Error
}

func (r *postRepository) GetByID(id uint) (*post.Post, error) {
	var post post.Post
	err := r.db.First(&post, id).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *postRepository) GetAll() ([]*post.Post, error) {
	var posts []*post.Post
	err := r.db.Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}
