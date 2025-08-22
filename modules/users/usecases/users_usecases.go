package usecases

import (
	"fiber-postgres-api/modules/entities/interfaces"
	"fiber-postgres-api/modules/models"
)

type usersUse struct {
	UsersRepo interfaces.UsersRepository
}

// Constructor
func NewUsersUsecase(usersRepo interfaces.UsersRepository) interfaces.UsersUsecase {
	return &usersUse{
		UsersRepo: usersRepo,
	}
}

func (u *usersUse) GetUserAndOrderListById(id string) (*models.GetUserAndOrderListByIdRes, error) {

	res, err := u.UsersRepo.GetUserAndOrderListById(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
