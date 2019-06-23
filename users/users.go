package users

import (
	"github.com/pkg/errors"

	"github.com/arturovm/epitome"
	"github.com/arturovm/epitome/storage"
)

// Users is a user managing service.
type Users struct {
	users storage.UserRepository
}

// ErrInvalidPassword is returned when the given password doesn't satisfy
// the minimum criteria.
var ErrInvalidPassword = errors.New("password is invalid")

// New takes a user repository and returns an initialized users service.
func New(users storage.UserRepository) *Users {
	return &Users{users: users}
}

// SignUp attempts to create a new user with the given username and password.
func (u *Users) SignUp(username, password string) error {
	user, err := epitome.NewUser(username, password)
	if err != nil {
		return errors.Wrap(err, "error creating new user")
	}

	err = u.users.Add(*user)
	if err != nil {
		return errors.Wrap(err, "error saving user")
	}
	return nil
}

// UserInfo retrieves a user instance from the database with the given username.
func (u *Users) UserInfo(username string) (*epitome.User, error) {
	return u.users.ByUsername(username)
}
