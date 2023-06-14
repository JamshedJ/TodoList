package service

import (
	"context"
	"github.com/sirupsen/logrus"
	"taskman4/pkg/repository"

	"taskman4/models"
)

type TaskService struct {
	repo repository.Tasks
}

func NewTaskService(repo repository.Tasks) *TaskService {
	return &TaskService{repo: repo}
}

func (ts *TaskService) GetTask(ctx context.Context, id, userID int) (task models.Task, err error) {
	if id <= 0 {
		return task, models.ErrBadRequest
	}
	task, err = ts.repo.GetTask(ctx, id, userID)
	if err != nil && err != models.ErrNoRows {
		logrus.Println("app GetTask", err)
	}
	return
}

func (ts *TaskService) GetAllTasks(ctx context.Context, userID int) (tasks []models.Task, err error) {
	tasks, err = ts.repo.GetAllTasks(ctx, userID)
	if err != nil && err != models.ErrNoRows {
		logrus.Println("app GetAllTasks", err)
	}
	return
}

func (ts *TaskService) AddTask(ctx context.Context, t models.Task) (id int, err error) {
	if !t.Validate() {
		return 0, models.ErrBadRequest
	}
	id, err = ts.repo.AddTask(ctx, t)
	if err != nil {
		logrus.Println("app AddTask", err)
	}
	return
}

func (ts *TaskService) UpdateTask(ctx context.Context, id, userID int, t models.Task) (err error) {
	if id <= 0 || !t.Validate() {
		return models.ErrBadRequest
	}
	err = ts.repo.UpdateTask(ctx, id, userID, t)
	if err != nil {
		logrus.Println("app UpdateTask", err)
	}
	return
}

func (ts *TaskService) DeleteTask(ctx context.Context, id, userID int) (err error) {
	if id <= 0 {
		return models.ErrBadRequest
	}
	err = ts.repo.DeleteTask(ctx, id, userID)
	if err != nil {
		logrus.Println("app DeleteTask", err)
	}
	return
}

func (ts*TaskService) MarkTask(ctx context.Context, id, userID int) (err error) {
	if id <= 0 {
		return models.ErrBadRequest
	}
	err = ts.repo.MarkTask(ctx, id, userID, false)
	if err != nil {
		logrus.Println("app MarkTask", err)
	}
	return
}

func (ts *TaskService) GetOverdueTasks(ctx context.Context, userID int) (tasks []models.Task, err error) {
	tasks, err = ts.repo.GetOverdueTasks(ctx)
	if err != nil && err != models.ErrNoRows {
		logrus.Println("app GetOverdueTasks", err)
	}
	return
}

func (ts *TaskService) ReassignTask(ctx context.Context, userID, taskID int, newUser string) (err error) {
	if len(newUser) < 3 || taskID <= 0 {
		return models.ErrBadRequest
	}
	err = ts.repo.ReassignTask(ctx, userID, taskID)
	if err != nil {
		logrus.Println("app ReassignTask", err)
	}
	return
}
