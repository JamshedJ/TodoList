package repository

import (
	"taskman4/models"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func newConnection(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})
	if err != nil {
		logrus.Fatal("Unable to connect to database: ", err)
	}
	if err = db.AutoMigrate(&models.User{}, &models.Task{}); err != nil {
		logrus.Println("Error migrating database: ", err)
	}
	return db
}

func closeConnection(db *gorm.DB, name string) {
	d, err := db.DB()
	if err != nil {
		logrus.Fatalf("Error getting %s instance: %v\n", name, err)
		return
	}
	if err = d.Close(); err != nil {
		logrus.Fatalf("Error closing %s: %v\n", name, err)
	}
}

// type DB struct {
// 	db *gorm.DB
// }

//
// func (d *DB) Close() {
// 	db, err := d.db.DB()
// 	if err != nil {
// 		log.Fatal("Error getting database instance: ", err)
// 	}
// 	if err = db.Close(); err != nil {
// 		log.Fatal("Error closing database: ", err)
// 	}
// }
