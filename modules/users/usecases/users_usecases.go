package usecases

import (
	"errors"
	"fiber-postgres-api/modules/entities"
	"regexp"
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

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func (u *usersUse) GetUserAndOrderListById(id string) (*entities.GetUserAndOrderListByIdRes, error) {
	if id == "" {
		return nil, errors.New("user id is required")
	}

	res, err := u.UsersRepo.GetUserAndOrderListById(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (u *usersUse) UserLogin(email string, password string) (*entities.UserLoginRes, error) {
	if email == "" || !emailRegex.MatchString(email) {
		return nil, errors.New("invalid email format")
	}
	if len(password) < 8 {
		return nil, errors.New("password must be at least 8 characters long")
	}

	// ดำเนินการ login
	res, err := u.UsersRepo.UserLogin(email, password)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (u *usersUse) UserLogin2(email string, password string) (*entities.UserLoginRes, error) {
	if email == "" || !emailRegex.MatchString(email) {
		return nil, errors.New("invalid email format")
	}
	if len(password) < 8 {
		return nil, errors.New("password must be at least 8 characters long")
	}

	// ดำเนินการ login
	res, err := u.UsersRepo.UserLogin(email, password)
	if err != nil {
		return nil, err
	}
	return res, nil
}
