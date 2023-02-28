package model

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Age int `json:"age"`
	Email string `json:"email"`
	Password string `json:"password"`
	IsPass bool `json:"is_pass"`
}