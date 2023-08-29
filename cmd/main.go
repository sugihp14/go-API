package main

import (
	"go-crud/pkg/infrastucture/database"
	"go-crud/pkg/initializers"
	"go-crud/pkg/migrations"
	"os"

	"go-crud/cmd/routers/post"
	"go-crud/cmd/routers/user"

	"github.com/gin-gonic/gin"
)

func main() {
	initializers.LoadEnvVariables()
	database.ConnectToDB()

	jwtSecret := []byte(os.Getenv("JWT_SECRET_KEY"))

	err := migrations.CreateUserTable(database.DB)
	if err != nil {
		panic("Failed to apply user table migration: " + err.Error())
	}
	err1 := migrations.CreatePostTable(database.DB)
	if err1 != nil {
		panic("Failed to apply Post table migration: " + err.Error())
	}

	r := gin.Default()

	post.SetupPostRouter(r, jwtSecret)
	user.SetupUserRouter(r, jwtSecret)

	port := os.Getenv("SERVER_PORT")
	r.Run(":" + port)
}
