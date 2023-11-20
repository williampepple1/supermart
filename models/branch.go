package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Branch struct {
	gorm.Model
	Name      string
	Location  string
	Manager   User      `gorm:"foreignKey:ManagerID"`
	ManagerID uuid.UUID `gorm:"type:uuid;not null" json:"manager_id"`
}
