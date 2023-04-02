package users

import (
	"github.com/gofiber/fiber/v2"

	"go-template/internal/model/entity"
	"go-template/internal/repository"
	"go-template/pkg/errors"
)

// Controller ...
type Controller struct {
}

// New ...
func New() Controller {
	return Controller{}
}

func (s Controller) GetAllUser(c *fiber.Ctx) error {
	users, err := repository.GetAllUser()
	if err != nil {
		return errors.Error(errors.NotFound, "error")
	}
	return c.JSON(users)
}

func (s Controller) CreateUser(c *fiber.Ctx) error {
	req := new(entity.CreateUser)
	if err := c.BodyParser(req); err != nil {
		return errors.Error(errors.InvalidArgument, "Please pass the correct argument")
	}

	user, err := repository.CreateUser(req)
	if err != nil {
		return errors.Error(errors.FailedPrecondition, "Create user failure")
	}

	return c.JSON(user)
}

func (s Controller) GetUserById(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := repository.GetUserById(id)
	if err != nil {
		return errors.Error(errors.NotFound, "Not found this userId")
	}

	return c.JSON(user)
}

func (s Controller) UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")

	req := new(entity.CreateUser)
	if err := c.BodyParser(req); err != nil {
		return errors.Error(errors.InvalidArgument, "Please pass the correct argument")
	}

	user, err := repository.UpdateUser(id, req)
	if err != nil {
		return errors.Error(errors.NotFound, "Not found this userId")
	}

	return c.JSON(user)
}

func (s Controller) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	err := repository.DeleteUser(id)
	if err != nil {
		return errors.Error(errors.NotFound, "Not found this userId")
	}

	return c.JSON("Xóa user thành công")
}
