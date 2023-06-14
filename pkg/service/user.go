package service

import (
	"context"
	"github.com/sirupsen/logrus"

	"taskman4/pkg/repository"

	"taskman4/models"
)

type UserService struct {
	repo repository.Users
}

func NewUserService(repo repository.Users) *UserService {
	return &UserService{repo: repo}
}

func (us *UserService) GetUser(ctx context.Context, userID int) (user models.User, err error) {
	user, err = us.repo.GetUser(ctx, userID)
	if err != nil && err != models.ErrNoRows {
		logrus.Println("app GetUser", err)
	}
	return
}

func (us *UserService) GetAllUsers(ctx context.Context) (users []models.User, err error) {
	users, err = us.repo.GetAllUsers(ctx)
	if err != nil && err != models.ErrNoRows {
		logrus.Println("app GetAllUsers", err)
	}
	return
}

func (us *UserService) AddUser(ctx context.Context, u models.User) (err error) {
	if !u.Validate() {
		err = models.ErrBadRequest
		return
	}
	u.Password = generatePasswordHash(u.Password)
	err = us.repo.AddUser(ctx, u)
	if err != nil {
		logrus.Println("app AddUser", err)
	}
	return
}

func (us *UserService) DeleteUser(ctx context.Context, userID int) (err error) {
	err = us.repo.DeleteUser(ctx, userID)
	if err != nil {
		logrus.Println("app DeleteUser", err)
	}
	return
}

func (us *UserService) UpdateUser(ctx context.Context, u models.User) (err error) {
	if !u.Validate() {
		return models.ErrBadRequest
	}
	u.Password = generatePasswordHash(u.Password)
	err = us.repo.UpdateUser(ctx, u)
	if err != nil {
		logrus.Println("app UpdateUser", err)
	}
	return
}
