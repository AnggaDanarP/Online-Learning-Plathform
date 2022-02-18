package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id        int    `json:"id" gorm:"primary_key"`
	Name      string `json:"name"`
	Email     string `json:"email" gorm:"unique_index"`
	Password  []byte `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}