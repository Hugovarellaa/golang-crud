package models

import "gorm.io/gorm"

type Customers struct {
	gorm.Model
	Id        uint64 `json:"id" gorm:"autoIncrement;primaryKey"`
	FirstName string `json:"first_name" gorm:"unique"`
	LastName  string `json:"last_name" gorm:"not null"`
	Email     string `json:"email" gorm:"not null"`
	Password  string `json:"password" gorm:"not null"`
}
