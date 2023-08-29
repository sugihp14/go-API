package post

type Service interface {
	CreatePost(post *Post) error
	UpdatePost(post *Post) error
	DeletePost(post *Post) error
	GetPostByID(id uint) (*Post, error)
	GetAllPosts() ([]*Post, error)
}
