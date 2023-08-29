package post

import (
	"go-crud/pkg/application/post"
	"go-crud/pkg/infrastucture/database"
	"go-crud/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func SetupPostRouter(r *gin.Engine, jwtSecret []byte) {
	postRepository := post.NewRepository(database.DB)
	postService := post.NewService(postRepository)
	postController := post.NewPostController(postService)

	authGroup := r.Group("/api/v1")
	authGroup.Use(middleware.JWTMiddleware(jwtSecret))

	authGroup.POST("/posts", postController.CreatePost)
	authGroup.GET("/posts", postController.GetAllPosts)
	authGroup.GET("/posts/:id", postController.GetPostByID)
	authGroup.PUT("/posts/:id", postController.UpdatePost)
	authGroup.DELETE("/posts/:id", postController.DeletePost)
}
