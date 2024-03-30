package repository

import (
	"errors"

	"github.com/vokhanh12/RongchoiApp/backend/entity"
)

var ErrLogin = errors.New("Video not found")

type AuthRepository interface {
	Login(username string, password string) (*entity.User, error)
	Register() (*entity.User, error)
	Logout() (*entity.User, error)
}
