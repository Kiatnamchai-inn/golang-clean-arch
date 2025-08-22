package interfaces

import "fiber-postgres-api/modules/models"

type UsersUsecase interface {
	GetUserAndOrderListById(id string) (*models.GetUserAndOrderListByIdRes, error)
}

type UsersRepository interface {
	GetUserAndOrderListById(id string) (*models.GetUserAndOrderListByIdRes, error)
}
