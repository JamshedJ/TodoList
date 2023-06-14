package service

import (
	"context"

	"taskman4/models"
	"taskman4/pkg/repository"
)

type Authorization interface {
	GenerateToken(ctx context.Context, u models.User) (string, error)
	ParseToken(jwtString string) (int, error)
}

type Tasks interface {
	GetTask(ctx context.Context, id, userID int) (task models.Task, err error)
	GetAllTasks(ctx context.Context, userID int) (tasks []models.Task, err error)
	AddTask(ctx context.Context, t models.Task) (id int, err error)
	UpdateTask(ctx context.Context, id, userID int, t models.Task) (err error)
	DeleteTask(ctx context.Context, id, userID int) (err error)
	MarkTask(ctx context.Context, id, userID int) (err error)
	GetOverdueTasks(ctx context.Context, userID int) (tasks []models.Task, err error)
	ReassignTask(ctx context.Context, userID, taskID int, newUser string) (err error)
}

type Users interface {
	GetUser(ctx context.Context, userID int) (user models.User, err error)
	GetAllUsers(ctx context.Context) (users []models.User, err error)
	AddUser(ctx context.Context, u models.User) (err error)
	DeleteUser(ctx context.Context, userID int) (err error)
	UpdateUser(ctx context.Context, u models.User) (err error)
}

type Service struct {
	Authorization
	Tasks
	Users
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Users),
		Tasks:         NewTaskService(repos.Tasks),
		Users:         NewUserService(repos.Users),
	}
}
