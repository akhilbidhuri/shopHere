package models

import (
	"time"

	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Id        int       `json:"id" gorm:"primary_key;auto_increment"`
	Name      string    `json:"name" gorm:"type:varchar(200)"`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
}

func (i *Item) Create(db *gorm.DB) error {
	err := db.Create(&i).Error
	if err != nil {
		return err
	}
	return nil
}

func (i *Item) GetAllItems(db *gorm.DB) (*[]Item, error) {
	items := []Item{}
	err := db.Model(&Item{}).Find(&items).Error
	if err != nil {
		return &[]Item{}, err
	}
	return &items, err
}

func (i *Item) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":         i.Id,
		"name":       i.Name,
		"created_at": i.CreatedAt.String(),
	}
}
