package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type RoleStatus string

const (
	StatusOwner   RoleStatus = "owner"
	StatusManager RoleStatus = "manager"
	StatusCashier RoleStatus = "cashier"
)

type User struct {
	ID        uuid.UUID  `gorm:"type:uuid;primary_key;" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	Username  string     `gorm:"unique" json:"username"`
	Password  string     `json:"password"`
	Role      RoleStatus `gorm:"type:varchar(20);not null;default:'cashier'" json:"role"`
	BranchID  uint       `json:"branch_id"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (base *User) BeforeCreate(tx *gorm.DB) (err error) {
	base.ID = uuid.New()
	currentTime := time.Now()
	base.CreatedAt = currentTime
	base.UpdatedAt = currentTime
	return nil
}

// GORM V2 uses callbacks like BeforeUpdate to handle the update timestamp
func (user *User) BeforeUpdate(tx *gorm.DB) (err error) {
	user.UpdatedAt = time.Now()
	return
}
