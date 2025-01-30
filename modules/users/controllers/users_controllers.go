package controllers

import (
	"fiber-postgres-api/modules/entities"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type usersController struct {
	UsersUse entities.UsersUsecase
}

func NewUsersController(r fiber.Router, usersUse entities.UsersUsecase) {
	controllers := &usersController{
		UsersUse: usersUse,
	}
	// slecet user and order list by user id using preloading and join condition
	// r.Get("/:userId", controllers.GetUserAndOrderListById)
	r.Post("/login", controllers.UserLogin)
	r.Get("/login", controllers.UserLogin2)
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

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":      "OK",
		"status_code": fiber.StatusOK,
		"message":     "get user data successfully.",
		"result":      res,
	})
}

func (h *usersController) UserLogin(c *fiber.Ctx) error {
	// Get email and password from request
	userLoginReq := entities.UserLoginReq{}
	if err := c.BodyParser(&userLoginReq); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":      http.StatusBadRequest,
			"status_code": http.StatusBadRequest,
			"message":     err.Error(),
			"result":      nil,
		})
	}

	res, err := h.UsersUse.UserLogin(userLoginReq.Email, userLoginReq.Password)
	if err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"status":      fiber.ErrInternalServerError.Message,
			"status_code": fiber.ErrInternalServerError.Code,
			"message":     err.Error(),
			"result":      nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":      "OK",
		"status_code": fiber.StatusOK,
		"message":     "login successfully.",
		"result":      res,
	})
}

func (h *usersController) UserLogin2(c *fiber.Ctx) error {
	// Get email and password from request
	email := c.Query("email")
	password := c.Query("password")

	res, err := h.UsersUse.UserLogin2(email, password)
	if err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"status":      fiber.ErrInternalServerError.Message,
			"status_code": fiber.ErrInternalServerError.Code,
			"message":     err.Error(),
			"result":      nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":      "OK",
		"status_code": fiber.StatusOK,
		"message":     "login successfully.",
		"result":      res,
	})
}
