package models

import (
	"errors"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage struct {
	db *gorm.DB
}

func (s *Storage) Setup() error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return errors.New(err.Error())
	}
	s.db = db
	s.db.AutoMigrate(&Cart{}, &User{}, &Item{}, &CartItem{}, &Order{})
	tx := s.db.Exec(`ALTER TABLE carts
	ADD FOREIGN KEY (user_id) 
	REFERENCES users(id)`)
	if tx.Error != nil {
		return errors.New(tx.Error.Error())
	}
	return nil
}
