package db

import (
	"fmt"
	"log"
	"nipatest/main/internal/model"

	"os"

	"github.com/jinzhu/gorm"
	// Required to import
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Initialize function that open db from configulation
func Initialize() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./ticket.db")
	if err != nil {
		log.Fatal(fmt.Sprintf("gorm open err: %s", err.Error()))
	}
	db.DB().SetMaxIdleConns(2)
	db.LogMode(true)
	return db
}

// TestDB func that create test db
func TestDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./../ticket_test.db")
	if err != nil {
		log.Fatal(fmt.Sprintf("test mode gorm open err: %s", err.Error()))
	}
	db.DB().SetMaxIdleConns(3)
	db.LogMode(false)
	return db
}

// DropTestDB func for del file db
func DropTestDB() error {
	if err := os.Remove("./../ticket_test.db"); err != nil {
		return err
	}
	return nil
}

// Paginate required page and pageSize if you want to customize page
func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

// AutoMigrate func for migrate db
func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&model.Ticket{},
	)
}
