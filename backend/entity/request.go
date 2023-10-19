package entity

import (


	"gorm.io/gorm"
)

type Request struct {

	gorm.Model

	image  string `gorm:"type:longtext"`

	grade string

	description string
	// FK 
	CustomerID *uint
	Customer Customer `gorm:"foreignKey:CustomerID"`

}