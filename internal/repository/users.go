package repository

import (
	"errors"
	"go-template/internal/model/entity"
	"time"
)

func GetAllUser() (users []entity.User, err error) {
	err = DB.Model(&users).Select()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func CreateUser(req *entity.CreateUser) (user *entity.User, err error) {
	if req.FullName == "" {
		return nil, errors.New("tên không được để trống")
	}

	user = &entity.User{
		Id:        NewID(),
		FullName:  req.FullName,
		Email:     req.Email,
		Phone:     req.Phone,
		CreatedAt: time.Now(),
	}
	_, err = DB.Model(user).WherePK().Returning("*").Insert()
	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetUserById(id string) (user *entity.User, err error) {
	user = &entity.User{
		Id: id,
	}
	err = DB.Model(user).WherePK().Select()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateUser(id string, req *entity.CreateUser) (user *entity.User, err error) {
	if req.FullName == "" {
		return nil, errors.New("tên không được để trống")
	}
	user = &entity.User{
		Id:        id,
		FullName:  req.FullName,
		Email:     req.Email,
		Phone:     req.Phone,
		UpdatedAt: time.Now(),
	}
	_, err = DB.Model(user).Column("full_name", "phone", "email", "updated_at").Returning("*").WherePK().Update()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func DeleteUser(id string) (err error) {
	user := entity.User{
		Id: id,
	}

	_, err = DB.Model(&user).WherePK().Delete()
	if err != nil {
		return err
	}

	return nil
}
