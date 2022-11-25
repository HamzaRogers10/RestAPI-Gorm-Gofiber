package models

import (
	"gorm.io/gorm"
)

type Developer struct {
	gorm.Model
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}
