package repository

import (
	"context"

	"taskman4/models"
)

type Users interface {
	Close()
	AddUser(ctx context.Context, u models.User) (err error)
	AuthenticateUser(ctx context.Context, u models.User) (id int, err error)
	GetUser(ctx context.Context, id int) (user models.User, err error)
	GetAllUsers(ctx context.Context) (users []models.User, err error)
	DeleteUser(ctx context.Context, id int) (err error)
	UpdateUser(ctx context.Context, u models.User) (err error)
}

type Tasks interface {
	Close()
	GetTask(ctx context.Context, id, userID int) (task models.Task, err error)
	GetAllTasks(ctx context.Context, userID int) (tasks []models.Task, err error)
	AddTask(ctx context.Context, t models.Task) (id int, err error)
	UpdateTask(ctx context.Context, id, userID int, t models.Task) (err error)
	DeleteTask(ctx context.Context, id, userID int) (err error)
	MarkTask(ctx context.Context, id, userID int, completed bool) (err error)
	GetOverdueTasks(ctx context.Context) (tasks []models.Task, err error)
	ReassignTask(ctx context.Context, id, userID int) (err error)
}

type Repository struct {
	Users
	Tasks
}

func NewRepository(cfg string) *Repository {
	userDB := newConnection(cfg)
	tasksDB := newConnection(cfg)
	return &Repository{
		Users: NewUserDB(userDB),
		Tasks: NewTaskDB(tasksDB),
	}
}

func (r *Repository) Close() {
	r.Users.Close()
	r.Tasks.Close()
}
