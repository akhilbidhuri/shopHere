package models

import "time"

type Cart struct {
	id           int       `json:"id" gorm:"primary_key"`
	user_id      int       `json:"user_id"`
	is_purchased bool      `json:"is_purchased"`
	created_at   time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
}
