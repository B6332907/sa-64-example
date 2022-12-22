package entity

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Description string
	Officer     []Officer `gorm:"foreignKey:Role_ID"`
}

type Gender struct {
	gorm.Model
	Description string
	Officer     []Officer `gorm:"foreignKey:Gender_ID"`
	Patiend     []Patiend `gorm:"foreignKey:Policing_ID"`
}

type Prefix struct {
	gorm.Model
	Description string
	Officer     []Officer `gorm:"foreignKey:Prefix_ID"`
	Patiend     []Patiend `gorm:"foreignKey:Policing_ID"`
}

type Policing struct {
	gorm.Model
	Description string
	Patiend     []Patiend `gorm:"foreignKey:Policing_ID"`
}

type Patiend struct {
	gorm.Model
	Name          string
	Date_of_Birth string
	Age           uint
	Address       string
	ID_Card       uint
	Phone         string

	Prefix_ID *uint
	Prefix    Prefix `gorm:"references:id"`

	Gender_ID *uint
	Gender    Gender `gorm:"references:id"`

	Policing_ID *uint
	Policing    Policing `gorm:"references:id"`
}

type Officer struct {
	gorm.Model
	Name     string
	Age      uint
	Phone    string
	Email    string
	Password string

	Prefix_ID *uint
	Prefix    Prefix `gorm:"references:id"`

	Gender_ID *uint
	Gender    Gender `gorm:"references:id"`

	Role_ID *uint
	Role    Role `gorm:"references:id"`
}
