package usecases

import (
	"errors"
	"fiber-postgres-api/modules/entities"
)

type usersUse struct {
	UsersRepo entities.UsersRepository
}

// Constructor
func NewUsersUsecase(usersRepo entities.UsersRepository) entities.UsersUsecase {
	return &usersUse{
		UsersRepo: usersRepo,
	}
}

func (u *usersUse) GetUserAndOrderListById(id string) (*entities.GetUserAndOrderListByIdRes, error) {
	if id == "" {
		return nil, errors.New("user id is required.")
	}

	res, err := u.UsersRepo.GetUserAndOrderListById(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
