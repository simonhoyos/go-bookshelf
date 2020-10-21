package model

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	Name string `gorm:"not null" json:"name"`
}
