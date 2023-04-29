package entity

import (
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Name     string `valid:"required~Name cannot be blank"`
	Email    string `valid:"required~Email cannot be blank, email"`
	Password string `valid:"required~Password cannot be blank"`
	Role     string

	Trade []Trade `gorm:"foreignKey:AdminID"`
	Sells []Sell  `gorm:"foreignKey:AdminID"`
	Drug  []Drug  `gorm:"foreignKey:AdminID"`
	Sales []Sales `gorm:"foreignKey:AdminID"`
}

type User struct {
	gorm.Model
	Name     string `valid:"required~Name cannot be blank"`
	Personal string `valid:"required,matches(^\\d{13}$)~Personal must be contain 13 numbers"`
	Email    string `valid:"required~Email cannot be blank, email"`
	Address  string `valid:"required~Address cannot be blank"`
	Password string `valid:"required~Password cannot be blank"`
	Tel      string `valid:"required,matches(^[0]\\d{9}$)~Tel must be contain 10 numbers"`
	BirthdayTime time.Time `valid:"timenotfuture~Birthday must be in the past"`
	Role     string

	// GenderID เป็น FK
	GenderID *uint
	// ข้อมูลของ Gender เมื่อ join ตาราง
	Gender Gender `gorm:"referenes:id" valid:"-"`

	// TitleID เป็น FK
	TitleID *uint
	// ข้อมูลของ Title เมื่อ join ตาราง
	Title Title `gorm:"referenes:id" valid:"-"`

	Drug  []Drug `gorm:"foreignKey:UserID"`
	Sells []Sell `gorm:"foreignKey:UserID"`
}

type Unit struct {
	gorm.Model
	Name string `valid:"required~Name cannot be blank"`

	Drug []Drug `gorm:"foreignKey:UnitID"`
}

type Drug struct {
	gorm.Model
	Code   string `valid:"required~Code cannot be blank"`
	Name   string `valid:"required~Name cannot be blank"`
	Amount int `valid:"required~Amount cannot be blank"`
	Price  int `valid:"required~Price cannot be blank"`

	// AdminID เป็น FK
	AdminID *uint
	// ข้อมูลของ Admin เมื่อ join ตาราง
	Admin Admin `gorm:"referenes:id" valid:"-"`

	// UserID เป็น FK
	UserID *uint
	// ข้อมูลของ User เมื่อ join ตาราง
	User User `gorm:"referenes:id" valid:"-"`

	// UnitID เป็น FK
	UnitID *uint
	// ข้อมูลของ Unit เมื่อ join ตาราง
	Unit Unit `gorm:"referenes:id" valid:"-"`

	Sells []Sell  `gorm:"foreignKey:DrugID"`
	Trade []Trade `gorm:"foreignKey:DrugID"`
}

type Sell struct {
	gorm.Model
	Quantity int `valid:"required~Quantity cannot be blank"`
	Cost     int `valid:"required~Cost cannot be blank"`
	Total    int `valid:"required~Total cannot be blank"`
	Payment  string `valid:"required~Payment cannot be blank"`
	SellTime time.Time `valid:"past~SelltTime not past"`

	// DrugID เป็น FK
	DrugID *uint
	// ข้อมูลของ Drug เมื่อ join ตาราง
	Drug Drug `gorm:"referenes:id" valid:"-"`

	// UnitID เป็น FK
	UnitID *uint
	// ข้อมูลของ Unit เมื่อ join ตาราง
	Unit Unit `gorm:"referenes:id" valid:"-"`

	// AdminID เป็น FK
	AdminID *uint
	// ข้อมูลของ Admin เมื่อ join ตาราง
	Admin Admin `gorm:"referenes:id" valid:"-"`

	// UserID เป็น FK
	UserID *uint
	// ข้อมูลของ User เมื่อ join ตาราง
	User User `gorm:"referenes:id" valid:"-"`

	// TradeID เป็น FK
	TradeID *uint
	// ข้อมูลของ Trade เมื่อ join ตาราง
	Trade Trade `gorm:"referenes:id" valid:"-"`

	Sales []Sales `gorm:"foreignKey:SellID"`
}

type Trade struct {
	gorm.Model
	QUANTITY  int `valid:"required~QUANTITY cannot be blank"`
	COST      int `valid:"required~COST cannot be blank"`
	Total     int `valid:"required~Total cannot be blank"`
	Evidence  string `valid:"required~Evidence cannot be blank"`
	TradeTime time.Time `valid:"past~TradeTime not past"`

	// AdminID เป็น FK
	AdminID *uint
	// ข้อมูลของ Admin เมื่อ join ตาราง
	Admin Admin `gorm:"referenes:id" valid:"-"`

	// UserID เป็น FK
	UserID *uint
	// ข้อมูลของ User เมื่อ join ตาราง
	User User `gorm:"referenes:id" valid:"-"`

	// DrugID เป็น FK
	DrugID *uint
	// ข้อมูลของ Drug เมื่อ join ตาราง
	Drug Drug `gorm:"referenes:id" valid:"-"`

	// StatusID เป็น FK
	StatusID *uint
	// ข้อมูลของ Status เมื่อ join ตาราง
	Status Status `gorm:"referenes:id" valid:"-"`

	Sell []Sell `gorm:"foreignKey:TradeID"`
	Sales []Sales `gorm:"foreignKey:TradeID"`
}

type Sales struct {
	gorm.Model
	DayTotal   int  `valid:"past~DayTotal not past"`
	MonthTotal int  `valid:"past~MonthTotal not past"`

	// AdminID เป็น FK
	AdminID *uint
	// ข้อมูลของ Admin เมื่อ join ตาราง
	Admin Admin `gorm:"referenes:id" valid:"-"`

	// SellID เป็น FK
	SellID *uint
	// ข้อมูลของ Sell เมื่อ join ตาราง
	Sell Sell `gorm:"referenes:id" valid:"-"`

	// TradeID เป็น FK
	TradeID *uint
	// ข้อมูลของ Trade เมื่อ join ตาราง
	Trade Trade `gorm:"referenes:id" valid:"-"`

}

type Status struct {
	gorm.Model
	Name string `valid:"required~Name cannot be blank"`

	Trade []Trade `gorm:"foreignKey:StatusID"`
}

type Title struct {
	gorm.Model
	Name string `valid:"required~Name cannot be blank"`

	Users []User `gorm:"foreignKey:TitleID"`
}
type Gender struct {
	gorm.Model
	Name string `valid:"required~Name cannot be blank"`

	Users []User `gorm:"foreignKey:GenderID"`
}