package post

import (
	"go-crud/pkg/domain/post"
)

type Service interface {
	CreatePost(title, body string) (*post.Post, error)
	UpdatePost(id uint, title, body string) (*post.Post, error)
	DeletePost(id uint) error
	GetPostByID(id uint) (*post.Post, error)
	GetAllPosts() ([]*post.Post, error)
}

type service struct {
	repository post.Repository
}

func NewService(repository post.Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) CreatePost(title, body string) (*post.Post, error) {
	newPost := &post.Post{
		Title: title,
		Body:  body,
	}
	if err := s.repository.Create(newPost); err != nil {
		return nil, err
	}
	return newPost, nil
}

func (s *service) UpdatePost(id uint, title, body string) (*post.Post, error) {
	postToUpdate, err := s.repository.GetByID(id)
	if err != nil {
		return nil, err
	}

	postToUpdate.Title = title
	postToUpdate.Body = body

	if err := s.repository.Update(postToUpdate); err != nil {
		return nil, err
	}
	return postToUpdate, nil
}

func (s *service) DeletePost(id uint) error {
	postToDelete, err := s.repository.GetByID(id)
	if err != nil {
		return err
	}
	return s.repository.Delete(postToDelete)
}

func (s *service) GetPostByID(id uint) (*post.Post, error) {
	return s.repository.GetByID(id)
}

func (s *service) GetAllPosts() ([]*post.Post, error) {
	return s.repository.GetAll()
}
