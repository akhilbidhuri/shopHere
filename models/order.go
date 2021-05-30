package models

import "time"

type Order struct {
	id         int       `json:"id" gorm:"primary_key`
	cart_id    int       `json:"cart_id" gorm:"type:integer REFERENCES carts(id)"` //ForeignKey(Cart)
	user_id    int       `json:"user_id" gorm:"type:integer REFERENCES users(id)"` //ForeignKey(User)
	created_at time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
}
