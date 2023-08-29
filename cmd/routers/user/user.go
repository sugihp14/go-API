package user

import (
	"go-crud/pkg/application/user"
	"go-crud/pkg/infrastucture/database"

	"github.com/gin-gonic/gin"
)

func SetupUserRouter(r *gin.Engine, jwtSecret []byte) {
	userService := user.NewUserService(database.NewUserRepository(database.DB), jwtSecret)
	userController := user.NewUserController(userService)

	r.POST("/register", userController.RegisterUser)
	r.POST("/login", userController.AuthenticateUser)
}
