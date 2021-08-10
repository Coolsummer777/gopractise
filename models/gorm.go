package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name     string
	Gender   string
	Age      int
	Email    string
	location string
}


