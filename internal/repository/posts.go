package repository

import (
	"errors"
	"go-template/internal/model/entity"
	"time"
)

func GetPostsOfUser(userId string) (posts []entity.Post, err error) {
	_, err = DB.Query(&posts, `
			SELECT p.id, p.title, p.content, p.author_id, p.created_at, p.updated_at
			FROM posts p
			INNER JOIN users u
			ON p.author_id = u.id
			WHERE u.id = ?
	`, userId)

	if err != nil {
		return nil, err
	}

	return posts, err
}

func CreatePost(userId string, req *entity.CreatePost) (post *entity.Post, err error) {
	// Kiểm tra title trong request có rỗng hay không?
	if req.Title == "" {
		return nil, errors.New("tiêu đề không được để trống")
	}

	// Tạo instance post dựa trên dữ liệu từ request
	post = &entity.Post{
		Id:        NewID(),
		Title:     req.Title,
		Content:   req.Content,
		AuthorId:  userId,
		CreatedAt: time.Now(),
	}

	// Insert vào trong database
	_, err = DB.Model(post).WherePK().Returning("*").Insert()
	if err != nil {
		return nil, err
	}

	return post, nil
}

func GetPostDetail(userId string, postId string) (post entity.Post, err error) {
	_, err = DB.Query(&post, `
			SELECT p.id, p.title, p.content, p.author_id, p.created_at, p.updated_at
			FROM posts p
			INNER JOIN users u
			ON p.author_id = u.id
			WHERE p.author_id = ? AND p.id = ?
	`, userId, postId)

	if err != nil {
		return entity.Post{}, err
	}

	return post, err
}

func UpdatePost(userId string, postId string, req *entity.CreatePost) (post *entity.Post, err error) {
	// Kiểm tra title trong request có rỗng hay không?
	if req.Title == "" {
		return nil, errors.New("tiêu đề không được để trống")
	}

	// Tạo instance post dựa trên dữ liệu từ request để tiến hành cập nhật dữ liệu
	post = &entity.Post{
		Id:        postId,
		Title:     req.Title,
		Content:   req.Content,
		AuthorId:  userId,
		UpdatedAt: time.Now(),
	}

	// Cập nhật các trường cần thiết trong database
	_, err = DB.Model(post).Column("title", "content", "updated_at").Returning("*").Where("id = ? AND author_id = ?", postId, userId).Update()
	if err != nil {
		return nil, err
	}

	return post, nil
}

func DeletePost(userId string, postId string) (err error) {
	_, err = DB.Exec(`
			DELETE FROM posts
			WHERE id = ? AND author_id = ?
	`, postId, userId)
	if err != nil {
		return err
	}

	return nil
}
