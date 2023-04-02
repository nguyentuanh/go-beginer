package homehandler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	"go-template/internal/app-api/base"
	"go-template/pkg/container"
)

// Controller ...
type Controller struct {
	base.Controller
}

// New ...
func New() Controller {
	obj := &Controller{}
	container.Fill(obj)
	return *obj
}

// Index godoc
// @Summary Index
// @Description Index api
// @Tags Home
// @Accept json
// @Produce json
// @Success 200 {object} response.CommonAPIResponse{} ok
// @Failure 400,404 {object} api.HTTPErrorResponse
// @Failure 500 {object} api.HTTPErrorResponse
// @Router /home/index [get]
func (s Controller) Index(c *fiber.Ctx) error {
	return s.Resp().WithMessage("success").WithStatus(http.StatusOK).Json(c)
}
