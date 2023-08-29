package models

import (
	"gorm.io/gorm"
)

type Price int

type Pizza struct {
	gorm.Model
	Name  string `gorm:"unique;not null"`
	Sauce string
	Price Price `gorm:"not null"`
}

type Crust struct {
	gorm.Model
	Name  string `gorm:"unique;not null"`
	Price Price  `gorm:"not null"`
}

type Size struct {
	gorm.Model
	Name       string  `json:"name" gorm:"unique;not null"`
	Multiplier float64 `json:"multiplier" gorm:"unique;not null"`
}

type Customer struct {
	gorm.Model
	Name   string `json:"name"`
	Number int64  `json:"number"`
}

type PizzaOrder struct {
	gorm.Model
	PizzaID    uint `gorm:"not null"`
	Pizza      Pizza
	CrustID    uint `gorm:"not null"`
	Crust      Crust
	SizeID     uint `gorm:"not null"`
	Size       Size
	CustomerID uint `gorm:"not null"`
	Customer   Customer
}

type Topping struct {
	gorm.Model
	Name       string `gorm:"unique;not null"`
	Price      Price  `gorm:"not null"`             // Price for topping when added to a medium sized pizza
	IsInternal bool   `json:"isInternal,omitempty"` // If true, than the user can not add this topping manually (but it can be used on pizza)
}

type PizzaOrderTopping struct {
	gorm.Model
	PizzaOrderID int        `gorm:"not null"`
	PizzaOrder   PizzaOrder `gorm:"constraint:OnDelete:CASCADE;"`
	ToppingID    int        `gorm:"not null"`
	Topping      Topping    `gorm:"constraint:OnDelete:CASCADE;"`
}
