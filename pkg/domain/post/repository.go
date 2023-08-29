package post

type Repository interface {
	Create(post *Post) error
	Update(post *Post) error
	Delete(post *Post) error
	GetByID(id uint) (*Post, error)
	GetAll() ([]*Post, error)
}
