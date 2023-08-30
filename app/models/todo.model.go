package models

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	Id         uint           `json:"id" gorm:"primaryKey; autoIncrement"`
	Title      string         `json:"title"`
	Completed  bool           `json:"completed" gorm:"default: false"`
	Created_at time.Time      `json:"created_at" gorm:"autoCreateTime"`
	Updated_at time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	Deleted_at gorm.DeletedAt `json:"deleted_at"`
}
