// pkg/domain/user_repository.go

package domain

type UserRepository interface {
    CreateUser(user *User) error
    GetUserByUsername(username string) (*User, error)
}
