// models/inventory.go

package models

import (
	"github.com/jinzhu/gorm"
)

type Inventory struct {
	gorm.Model
	ProductID uint
	Quantity  int
	BranchID  uint
}
