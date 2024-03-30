package auth

import "github.com/vokhanh12/RongchoiApp/backend/entity"

type UseCase interface {
	Login(username string, password string) (*entity.User, error)

	Register() (*entity.User, error)
	Logout() (*entity.User, error)
}
