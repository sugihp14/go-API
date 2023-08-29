package user

import (
	"go-crud/pkg/application/response"
	domain "go-crud/pkg/domain/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService domain.UserService
}

func NewUserController(userService domain.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) RegisterUser(c *gin.Context) {
	var registrationData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&registrationData); err != nil {
		response.ErrorJSON(c, http.StatusInternalServerError, "invalid Data")
		return
	}

	if err := uc.userService.RegisterUser(registrationData.Username, registrationData.Password); err != nil {
		response.ErrorJSON(c, http.StatusInternalServerError, "Failed to Register")
		return
	}
	response.JSON(c, http.StatusOK, gin.H{"message": "User registered successfully"})


}

func (uc *UserController) AuthenticateUser(c *gin.Context) {
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&loginData); err != nil {
		response.ErrorJSON(c, http.StatusInternalServerError, "invalid Data")
		return
	}

	token, err := uc.userService.AuthenticateUser(loginData.Username, loginData.Password)
	if err != nil {
		response.ErrorJSON(c, http.StatusInternalServerError, "invalid Credentials")
		return
	}

	response.JSON(c, http.StatusOK, gin.H{"token":token})
}
