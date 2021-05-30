package models

import "github.com/golang/protobuf/ptypes/timestamp"

type Item struct {
	id         int                 `json:"id" gorm:"primary_key"`
	name       string              `json:"name"`
	created_at timestamp.Timestamp `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
}
