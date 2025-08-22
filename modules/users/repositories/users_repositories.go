package repositories

import (
	"fiber-postgres-api/modules/entities"
	"fiber-postgres-api/modules/entities/interfaces"
	"fiber-postgres-api/modules/models"

	"gorm.io/gorm"
)

type usersRepo struct {
	Db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) interfaces.UsersRepository {
	return &usersRepo{
		Db: db,
	}
}

func (r *usersRepo) GetUserAndOrderListById(id string) (*models.GetUserAndOrderListByIdRes, error) {
	usersModel := entities.User{}
	res := models.GetUserAndOrderListByIdRes{}

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
