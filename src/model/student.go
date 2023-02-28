package model

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Age int `json:"age"`
	IsPass bool `json:"is_pass"`
}