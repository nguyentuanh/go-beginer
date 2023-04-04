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

func (s Controller) GetPostsOfUser(c *fiber.Ctx) error {
	// Đọc id của user trên URL
	userId := c.Params("id")
	posts, err := repository.GetPostsOfUser(userId)
	if err != nil {
		return errors.Error(errors.NotFound, "error")
	}
	return c.JSON(posts)
}

func (s Controller) CreatePost(c *fiber.Ctx) error {
	// Đọc id của user trên URL
	userId := c.Params("id")

	// Đọc dữ liệu từ body của request
	req := new(entity.CreatePost)
	if err := c.BodyParser(req); err != nil {
		return errors.Error(errors.FailedPrecondition, "Invali data post")
	}

	// Tạo post mới dựa trên uesrId và thông tin dữ liệu đọc được từ request
	post, err := repository.CreatePost(userId, req)
	if err != nil {
		return errors.Error(errors.FailedPrecondition, "Create new post failure")
	}

	// Trả về post mới sau khi tạo thành công cho client
	return c.JSON(post)
}

func (s Controller) GetPostDetail(c *fiber.Ctx) error {
	// Đọc id của user trên URL
	userId := c.Params("id")

	// Đọc id của post trên URL
	postId := c.Params("postId")

	// Lấy thông tin của post dựa trên userId và postId
	post, err := repository.GetPostDetail(userId, postId)
	if err != nil {
		return errors.Error(errors.NotFound, "Not found this Post")
	}

	// Trả thông tin post về cho client
	return c.JSON(post)
}

func (s Controller) UpdatePost(c *fiber.Ctx) error {
	// Đọc id của user trên URL
	userId := c.Params("id")

	// Đọc id của post trên URL
	postId := c.Params("postId")

	// Đọc dữ liệu từ body của request
	req := new(entity.CreatePost)
	if err := c.BodyParser(req); err != nil {
		return errors.Error(errors.FailedPrecondition, "Invali data post")
	}

	// Cập nhật thông tin post dựa trên uesrId và thông tin dữ liệu đọc được từ request
	post, err := repository.UpdatePost(userId, postId, req)
	if err != nil {
		return errors.Error(errors.FailedPrecondition, "Update post failure")
	}

	// Trả về post sau khi cập nhật thành công cho client
	return c.JSON(post)
}

func (s Controller) DeletePost(c *fiber.Ctx) error {
	// Đọc id của user trên URL
	userId := c.Params("id")

	// Đọc id của post trên URL
	postId := c.Params("postId")

	// Xóa post dựa trên uesrId và postId
	err := repository.DeletePost(userId, postId)
	if err != nil {
		return errors.Error(errors.NotFound, "Not found this postId")
	}

	return c.JSON("Xóa post thành công")
}
