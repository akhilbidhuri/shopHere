package models

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model
	Cart_id int `json:"cart_id" gorm:"not null"`
	Item_id int `json:"item_id" gorm:"not null"`
}

func (ci *CartItem) Create(db *gorm.DB) error {
	err := db.Debug().Create(&ci).Error
	if err != nil {
		return err
	}
	return nil
}

func (ci *CartItem) GetItemsForCart(db *gorm.DB) (*[]CartItem, error) {
	cartItems := []CartItem{}
	err := db.Debug().Model(&CartItem{}).Where("cart_id = ?", ci.Cart_id).Find(&cartItems).Error
	if err != nil {
		return &[]CartItem{}, err
	}
	return &cartItems, err
}

func (ci *CartItem) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"cart_id": ci.Cart_id,
		"item_id": ci.Item_id,
	}
}
