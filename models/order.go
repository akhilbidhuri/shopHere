package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Id        int       `json:"id" gorm:"primary_key;auto_increment"`
	Cart_id   int       `json:"cart_id" gorm:"type:integer REFERENCES carts(id)"` //ForeignKey(Cart)
	User_id   int       `json:"user_id" gorm:"type:integer REFERENCES users(id)"` //ForeignKey(User)
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
}

func (o *Order) Create(db *gorm.DB) error {
	err := db.Debug().Create(&o).Error
	if err != nil {
		return nil
	}
	return nil
}

func (o *Order) GetOrderByUID(db *gorm.DB) (*[]Order, error) {
	orders := []Order{}
	err := db.Debug().Model(&CartItem{}).Find(&orders).Where("user_id = ?", o.User_id).Error
	if err != nil {
		return &[]Order{}, err
	}
	return &orders, err
}

func (o *Order) GetAllOrders(db *gorm.DB) (*[]Order, error) {
	orders := []Order{}
	err := db.Debug().Model(&Order{}).Find(&orders).Error
	if err != nil {
		return &[]Order{}, err
	}
	return &orders, err
}

func (o *Order) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":         o.Id,
		"cart_id":    o.Cart_id,
		"user_id":    o.User_id,
		"created_at": o.CreatedAt.String(),
	}
}
