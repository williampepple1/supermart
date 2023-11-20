package models

import (
	"github.com/jinzhu/gorm"
)

type Branch struct {
	gorm.Model
	Name      string
	Location  string
	Manager   User `gorm:"foreignKey:ManagerID"`
	ManagerID uint
}
