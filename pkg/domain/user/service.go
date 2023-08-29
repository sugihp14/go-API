package domain

type UserService interface {
	RegisterUser(username, password string) error
	AuthenticateUser(username, password string) (string, error)
}
