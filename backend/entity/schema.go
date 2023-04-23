package entity

import (
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Role     string

	Trade []Trade `gorm:"foreignKey:AdminID"`
	Sells []Sell  `gorm:"foreignKey:AdminID"`
	Drug  []Drug  `gorm:"foreignKey:AdminID"`
	Sales []Sales `gorm:"foreignKey:AdminID"`
}

type User struct {
	gorm.Model
	Title    string
	Gender   string
	Name     string
	Personal string
	Email    string
	Address  string
	Password string
	Tel      string
	Role     string

	Drug  []Drug `gorm:"foreignKey:UserID"`
	Sells []Sell `gorm:"foreignKey:UserID"`
}

type Drug struct {
	gorm.Model
	Unit   string
	Code   string
	Name   string
	Amount int
	Price  int

	// AdminID เป็น FK
	AdminID *uint
	// ข้อมูลของ Admin เมื่อ join ตาราง
	Admin Admin

	// UserID เป็น FK
	UserID *uint
	// ข้อมูลของ User เมื่อ join ตาราง
	User User

	Sells []Sell  `gorm:"foreignKey:DrugID"`
	Trade []Trade `gorm:"foreignKey:DrugID"`
}

type Sell struct {
	gorm.Model
	Quantity string
	Cost     int
	Total    int
	Type     string
	Payment  string
	Status   string
	SellTime time.Time

	// DrugID เป็น FK
	DrugID *uint
	// ข้อมูลของ Drug เมื่อ join ตาราง
	Drug Drug

	// AdminID เป็น FK
	AdminID *uint
	// ข้อมูลของ Admin เมื่อ join ตาราง
	Admin Admin

	// UserID เป็น FK
	UserID *uint
	// ข้อมูลของ User เมื่อ join ตาราง
	User User

	// TradeID เป็น FK
	TradeID *uint
	// ข้อมูลของ Trade เมื่อ join ตาราง
	Trade Trade

	Sales []Sales `gorm:"foreignKey:SellID"`
}

type Trade struct {
	gorm.Model
	QUANTITY  string
	COST      string
	Total     string
	Type      string
	Status    string
	Evidence  string
	TradeTime time.Time

	// AdminID เป็น FK
	AdminID *uint
	// ข้อมูลของ Admin เมื่อ join ตาราง
	Admin Admin

	// UserID เป็น FK
	UserID *uint
	// ข้อมูลของ User เมื่อ join ตาราง
	User User

	// DrugID เป็น FK
	DrugID *uint
	// ข้อมูลของ Drug เมื่อ join ตาราง
	Drug Drug

	Sell []Sell `gorm:"foreignKey:TradeID"`
}

type Sales struct {
	gorm.Model
	DayTotal   int
	MonthTotal int

	// AdminID เป็น FK
	AdminID *uint
	// ข้อมูลของ Admin เมื่อ join ตาราง
	Admin Admin

	// SellID เป็น FK
	SellID *uint
	// ข้อมูลของ Sell เมื่อ join ตาราง
	Sell Sell

	// TradeID เป็น FK
	TradeID *uint
	// ข้อมูลของ Trade เมื่อ join ตาราง
	Trade Trade
}
