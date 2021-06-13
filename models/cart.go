package models

import (
	"time"

	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	Id           int       `json:"id" gorm:"primary_key;auto_increment"`
	User_id      int       `json:"user_id" gorm:"type:integer"`
	Is_purchased bool      `json:"is_purchased" gorm:"type:boolean;not null"`
	CreatedAt    time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
}

func (c *Cart) Create(db *gorm.DB) error {
	err := db.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&c).Error

		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *Cart) GetCartByID(db *gorm.DB, id int) error {
	err := db.Model(Cart{}).Where("id = ?", id).Take(&c).Error
	if err != nil {
		return err
	}
	return err
}

func (c *Cart) GetNewCart(db *gorm.DB) (*User, error) {
	newCart := Cart{}
	newCart.User_id = c.User_id
	newCart.Is_purchased = false
	if err := newCart.Create(db); err != nil {
		return &User{}, err
	}
	user := User{}
	if err := user.GetUserByID(db, c.User_id); err != nil {
		return &User{}, err
	}
	user.Cart_id = newCart.Id
	if err := user.Update(db); err != nil {
		return &User{}, err
	}
	return &user, nil
}

func (c *Cart) Update(db *gorm.DB) error {
	err := db.Save(&c).Error
	if err != nil {
		return nil
	}
	return nil
}

func (c *Cart) GetAllCarts(db *gorm.DB) (*[]Cart, error) {
	carts := []Cart{}
	err := db.Model(&Cart{}).Find(&carts).Error
	if err != nil {
		return &[]Cart{}, err
	}
	return &carts, err
}

func (c *Cart) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":           c.Id,
		"user_id":      c.User_id,
		"is_purchased": c.Is_purchased,
		"created_at":   c.CreatedAt.String(),
	}
}
