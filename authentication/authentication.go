package authentication

import (
	"github.com/pkg/errors"

	"github.com/arturovm/epitome"
	"github.com/arturovm/epitome/storage"
)

// ErrWrongCredentials is returned when the given username or passsword don't
// match those registered with the server.
var ErrWrongCredentials = errors.New("wrong username or password")

// Authentication is an authentication management service.
type Authentication struct {
	sessions storage.SessionRepository
	users    storage.UserRepository
}

// New takes a sessions repository and a users repository and returns an
// initialized authentication service.
func New(sessions storage.SessionRepository, users storage.UserRepository) *Authentication {
	return &Authentication{
		sessions: sessions,
		users:    users,
	}
}

// LogIn creates a new session if the given username and password match those
// registered with the server.
func (a *Authentication) LogIn(username, password string) (*epitome.Session, error) {
	_, err := a.users.ByUsername(username)
	if err != nil {
		return nil, errors.Wrap(err, "error retrieving user")
	}

	session, err := epitome.NewSession(username)
	if err != nil {
		return nil, errors.Wrap(err, "error creating session")
	}

	err = a.sessions.Add(*session)
	if err != nil {
		return nil, errors.Wrap(err, "error saving session")
	}

	return session, nil
}