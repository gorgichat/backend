package database

import (
	"fmt"
	"time"
	"log/slog"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
	ID       uint      `json:"id" gorm:"primaryKey"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func GetDB() *gorm.DB {
	if db == nil {
		return nil
	}
	return db
}

func ConnectDatabase(host string, port int, user, password, dbName string) error {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbName,
	)
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		slog.Error("database error", "Database connection failed", err.Error())
		return err
	}
	slog.Info("database", "Database connection established", "mysql_connected")
	return nil
}

func MigrateDatabase(db *gorm.DB) {
	if db := GetDB(); db != nil {
		_ = db.AutoMigrate(&User{})
	} else {
		slog.Error("database error", "Database connection failed", "db is nil")
	}
}