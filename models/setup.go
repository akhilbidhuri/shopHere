package models

import (
	"errors"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage struct {
	DB *gorm.DB
}

func (s *Storage) Setup() error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return errors.New(err.Error())
	}
	s.DB = db
	if err = s.DB.Migrator().AutoMigrate(&Cart{}, &User{}, &Item{}, &CartItem{}, &Order{}); err != nil {
		return err
	}
	if err = s.applyForeignKeyConstraints(); err != nil {
		return err
	}
	log.Println("DB migrated and setup done.")
	return nil
}

func (s *Storage) applyForeignKeyConstraints() error {
	tx1 := s.DB.Exec(`ALTER TABLE users
		ADD FOREIGN KEY (cart_id)
		REFERENCES carts(id)	
		DEFERRABLE INITIALLY DEFERRED
	`)
	if tx1.Error != nil {
		return errors.New(tx1.Error.Error())
	}
	tx2 := s.DB.Exec(`ALTER TABLE carts
		ADD FOREIGN KEY (user_id)
		REFERENCES users(id)	
		DEFERRABLE INITIALLY DEFERRED
	`)
	if tx2.Error != nil {
		return errors.New(tx2.Error.Error())
	}
	return nil
}
