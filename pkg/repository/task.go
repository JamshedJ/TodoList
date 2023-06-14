package repository

import (
	"context"
	"time"

	"taskman4/models"

	"gorm.io/gorm"
)

type TaskDB struct {
	db *gorm.DB
}

func NewTaskDB(db *gorm.DB) *TaskDB {
	return &TaskDB{
		db: db,
	}
}

func (d *TaskDB) Close() {
	closeConnection(d.db, "TaskDB")
}

func (d *TaskDB) GetTask(ctx context.Context, id, userID int) (task models.Task, err error) {
	err = d.db.First(&task, "id = ? AND user_id = ?", id, userID).Error
	if err == gorm.ErrRecordNotFound {
		err = models.ErrNoRows
	}
	return
}

func (d *TaskDB) GetAllTasks(ctx context.Context, userID int) (tasks []models.Task, err error) {
	err = d.db.Find(&tasks, "user_id = ?", userID).Error
	if err == gorm.ErrRecordNotFound {
		err = models.ErrNoRows
	}
	return
}

func (d *TaskDB) AddTask(ctx context.Context, t models.Task) (id int, err error) {
	err = d.db.Create(&t).Error
	if err == nil {
		id = t.ID
	}
	return
}

func (d *TaskDB) UpdateTask(ctx context.Context, id, userID int, t models.Task) (err error) {
	err = d.db.Where("id = ? AND user_id = ?", id, userID).Updates(&t).Error
	if err == gorm.ErrRecordNotFound {
		err = models.ErrNoRows
	}
	return
}

func (d *TaskDB) DeleteTask(ctx context.Context, id, userID int) (err error) {
	err = d.db.Where("id = ? AND user_id = ?", id, userID).Delete(&models.Task{}).Error
	if err == gorm.ErrRecordNotFound {
		err = models.ErrNoRows
	}
	return
}

func (d *TaskDB) MarkTask(ctx context.Context, id, userID int, completed bool) (err error) {
	err = d.db.Model(&models.Task{}).Where("id = ? AND user_id = ?", id, userID).Update("completed", completed).Error
	if err == gorm.ErrRecordNotFound {
		err = models.ErrNoRows
	}
	return
}

func (d *TaskDB) GetOverdueTasks(ctx context.Context) (tasks []models.Task, err error) {
	err = d.db.Find(&tasks, "deadline < ?", time.Now()).Error
	if err == gorm.ErrRecordNotFound {
		err = models.ErrNoRows
	}
	return
}

// func (d *TaskDB) GetOverdueTasks(ctx context.Context) (tasks []models.Task, err error) {
// 	err = d.db.Find(&tasks, "deadline < ? AND completed = ?", time.Now(), false).Error
// 	if err == gorm.ErrRecordNotFound {
// 		err = models.ErrNoRows
// 	}
// 	return                                     //column "completed" does not exist in database        
// }

func (d *TaskDB) ReassignTask(ctx context.Context, id, userID int) (err error) {
	err = d.db.Model(&models.Task{}).Where("id = ?", id).Update("user_id", userID).Error
	if err == gorm.ErrRecordNotFound {
		err = models.ErrNoRows
	}
	return
}
