package database

import (
	"go-crud/pkg/domain/post"

	"gorm.io/gorm"
)

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) post.Repository {
	return &postRepository{
		db: db,
	}
}

func (r *postRepository) Create(newPost *post.Post) error {
	return r.db.Create(newPost).Error
}

func (r *postRepository) Update(postToUpdate *post.Post) error {
	return r.db.Save(postToUpdate).Error
}

func (r *postRepository) Delete(postToDelete *post.Post) error {
	return r.db.Delete(postToDelete).Error
}

func (r *postRepository) GetByID(id uint) (*post.Post, error) {
	var postByID post.Post
	err := r.db.First(&postByID, id).Error
	if err != nil {
		return nil, err
	}
	return &postByID, nil
}

func (r *postRepository) GetAll() ([]*post.Post, error) {
	var posts []*post.Post
	err := r.db.Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}
