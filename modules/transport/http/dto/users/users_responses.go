package usersdto

import "fiber-postgres-api/modules/models"

func MapGetUserAndOrderListByIdRespDTO(res *models.GetUserAndOrderListByIdRes) *GetUserAndOrderListByIdResDTO {
	if res == nil {
		return nil
	}

	return &GetUserAndOrderListByIdResDTO{
		ID:     res.ID,
		Name:   res.Name,
		Email:  res.Email,
		Orders: res.Orders,
	}
}

type GetUserAndOrderListByIdResDTO struct {
	ID     int32               `json:"id"`
	Name   string              `json:"name"`
	Email  string              `json:"email"`
	Orders []models.UserOrders `json:"orders"`
}
