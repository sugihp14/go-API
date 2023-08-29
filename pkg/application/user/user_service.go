package user

import (
	"time"

	domain "go-crud/pkg/domain/user"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepository domain.UserRepository
	jwtSecret      []byte
}

func NewUserService(userRepository domain.UserRepository, jwtSecret []byte) domain.UserService {
	return &userService{
		userRepository: userRepository,
		jwtSecret:      jwtSecret,
	}
}

func (s *userService) RegisterUser(username, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	newUser := &domain.User{
		Username: username,
		Password: string(hashedPassword),
	}

	return s.userRepository.CreateUser(newUser)
}

func (s *userService) AuthenticateUser(username, password string) (string, error) {
	foundUser, err := s.userRepository.GetUserByUsername(username)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(password)); err != nil {
		return "", err
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": foundUser.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token berlaku selama 1 hari
	})

	tokenString, err := token.SignedString(s.jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
