package repositories

import (
	"fiber-postgres-api/modules/entities"
	"fiber-postgres-api/modules/entities/dbmodels"

	"gorm.io/gorm"
)

type usersRepo struct {
	Db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) entities.UsersRepository {
	return &usersRepo{
		Db: db,
	}
}

func (r *usersRepo) GetUserAndOrderListById(id string) (*entities.GetUserAndOrderListByIdRes, error) {
	usersModel := dbmodels.User{}
	res := entities.GetUserAndOrderListByIdRes{}

	err := r.Db.Model(&usersModel).
		Preload("Orders", func(db *gorm.DB) *gorm.DB {
			return db.Select("orders.*, products.name AS product_name, products.price AS product_price").
				Joins("LEFT JOIN products ON products.id = orders.product_id")
		}).
		Where("users.id = ?", id).
		First(&res).Error

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *usersRepo) UserLogin(email string, password string) (*entities.UserLoginRes, error) {
	usersModel := dbmodels.User{}
	res := entities.UserLoginRes{}

	err := r.Db.Model(&usersModel).
		Where("email = ? AND password = ?", email, password).
		First(&res).Error

	if err != nil {
		return nil, err
	}

	return &res, nil
}
