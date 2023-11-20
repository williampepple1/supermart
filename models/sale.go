package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Sale struct {
	gorm.Model
	ProductID uint
	Quantity  int
	Total     float64
	Date      time.Time
	BranchID  uint
}
