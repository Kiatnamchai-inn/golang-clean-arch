package controllers

import (
	"fiber-postgres-api/modules/entities/interfaces"
	entities "fiber-postgres-api/modules/entities/interfaces"
	usersdto "fiber-postgres-api/modules/transport/http/dto/users"

	"github.com/gofiber/fiber/v2"
)

type usersController struct {
	UsersUse entities.UsersUsecase
}

func NewUsersController(r fiber.Router, usersUse interfaces.UsersUsecase) {
	controllers := &usersController{
		UsersUse: usersUse,
	}
	// slecet user and order list by user id using preloading and join condition
	r.Get("/:userId", controllers.GetUserAndOrderListById)
}

func (h *usersController) GetUserAndOrderListById(c *fiber.Ctx) error {

	// Get user id from request
	userId := c.Params("userId")

	res, err := h.UsersUse.GetUserAndOrderListById(userId)
	if err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"status":      fiber.ErrInternalServerError.Message,
			"status_code": fiber.ErrInternalServerError.Code,
			"message":     err.Error(),
			"result":      nil,
		})
	}

	dto := usersdto.MapGetUserAndOrderListByIdRespDTO(res)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":      "OK",
		"status_code": fiber.StatusOK,
		"message":     "get user data successfully.",
		"result":      dto,
	})
}
