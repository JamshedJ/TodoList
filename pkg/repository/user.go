package repository

import (
	"context"

	"taskman4/models"

	"gorm.io/gorm"
)

type UserDB struct {
	db *gorm.DB
}

func NewUserDB(db *gorm.DB) *UserDB {
	return &UserDB{
		db: db,
	}
}

func (d *UserDB) Close() {
	closeConnection(d.db, "UserDB")
}

func (d *UserDB) AddUser(ctx context.Context, u models.User) (err error) {
	if err = d.db.WithContext(ctx).Create(&u).Error; err == gorm.ErrDuplicatedKey {
		err = models.ErrDuplicate
	}
	return
}

func (d *UserDB) AuthenticateUser(ctx context.Context, u models.User) (id int, err error) {
	if err = d.db.WithContext(ctx).Table("users").Select("id").
		Take(&id, "username = ? AND password = ?", u.Username, u.Password).Error; err == gorm.ErrRecordNotFound {
		err = models.ErrNoRows
	}
	return
}

func (d *UserDB) GetUser(ctx context.Context, id int) (user models.User, err error) {
	if err = d.db.WithContext(ctx).Take(&user, "id = ?", id).Error; err == gorm.ErrRecordNotFound {
		err = models.ErrNoRows
	}
	return
}

func (d *UserDB) GetAllUsers(ctx context.Context) (users []models.User, err error) {
	if err = d.db.WithContext(ctx).Select("username").Find(&users).Error; err == gorm.ErrRecordNotFound {
		err = models.ErrNoRows
	}
	return
}

func (d *UserDB) DeleteUser(ctx context.Context, id int) (err error) {
	if err = d.db.WithContext(ctx).Delete(&models.User{}, id).Error; err == gorm.ErrRecordNotFound {
		err = models.ErrNoRows
	}
	return
}

func (d *UserDB) UpdateUser(ctx context.Context, u models.User) (err error) {
	if err = d.db.WithContext(ctx).Model(&u).Updates(u).Error; err == gorm.ErrRecordNotFound {
		err = models.ErrNoRows
	}
	return
}
