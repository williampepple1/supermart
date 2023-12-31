// models/product.go

package models

import (
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Name        string
	Description string
	Price       float64
	Quantity    int
	BranchID    uint
}
