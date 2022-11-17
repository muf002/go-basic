package models

import (
	"time"

	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	Name       string
	Age        uint
	Gender     string
	Email      string
	Birthday   time.Time
	Department string
}
