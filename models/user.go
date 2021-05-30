package models

import "time"

type User struct {
	id         int       `json:"id" gorm:"primary_key;auto_increment;not null" validate:"required"`
	name       string    `json:"name" validate:"required"`
	username   string    `json:"username" validate:"required"`
	password   string    `json:"password" validate:"required"`
	token      string    `json:"token"`
	created_at time.Time `json:"create_at" gorm:"default:CURRENT_TIMESTAMP"`
	cart_id    int       `json:"cart_id" gorm:"type:integer REFERENCES carts(id)"`
	Cart       Cart      `gorm:"foreignkey:cart_id,constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
