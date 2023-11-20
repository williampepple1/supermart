// models/user.go

package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password string
	Role     string // "cashier", "manager", "owner"
	BranchID uint
}
