package auth

import (
	"errors"
	"fmt"

	"github.com/vokhanh12/RongchoiApp/backend/app/repository"
	"github.com/vokhanh12/RongchoiApp/backend/entity"
)

type usecase struct {
	auth_repo repository.AuthRepository
}

// Login implements UseCase.
func (uc *usecase) Login(username string, password string) (*entity.User, error) {
	data, err := uc.auth_repo.Login(username, password)
	if err != nil {
		if errors.Is(err, repository.ErrLogin) {
			return nil, ErrLogin
		}
		return nil, fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}
	return data, nil
}

// Logout implements UseCase.
func (*usecase) Logout() (*entity.User, error) {
	panic("unimplemented")
}

// Register implements UseCase.
func (*usecase) Register() (*entity.User, error) {
	panic("unimplemented")
}

func NewUseCase(auth_repo repository.AuthRepository) UseCase {
	return &usecase{auth_repo: auth_repo}
}
