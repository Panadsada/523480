package entity

import (
	//   "fmt"
	//   "time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {

	database, err := gorm.Open(sqlite.Open("sa-65.db"), &gorm.Config{})

	if err != nil {

		panic("failed to connect database")

	}

	// Migrate the schema
	database.AutoMigrate(
		&Drug{}, &Sell{}, &User{}, &Sales{}, &Admin{}, &Trade{},
	)

	db = database
	/////////////////////////////////////////
	password, err := bcrypt.GenerateFromPassword([]byte("123456"), 14) //password mail
	password1, err := bcrypt.GenerateFromPassword([]byte("555555"), 14)

	db.Model(&User{}).Create(&User{
		Name:     "Panadda Srisawat",
		Email:    "panadda@gmail.com",
		Password: string(password),
		Role:     "user",
	})

	db.Model(&Admin{}).Create(&Admin{
		Name:     "abc",
		Email:    "abc@gmail.com",
		Password: string(password1),
		Role:     "admin",
	})

	var panadda User
	var abc Admin

	db.Raw("SELECT * FROM users WHERE email = ?", "panadda@gmail.com").Scan(&panadda)
	db.Raw("SELECT * FROM admins WHERE email = ?", "abc@gmail.com").Scan(&abc)

	//---Title Data
	Mr := Title{
		Name: "นาย",
	}
	db.Model(&Title{}).Create(&Mr)

	Mrs := Title{
		Name: "นาง",
	}
	db.Model(&Title{}).Create(&Mrs)

	Miss := Title{
		Name: "นางสาว",
	}
	db.Model(&Title{}).Create(&Miss)

	//---Gender Data
	Male := Gender{
		Name: "ชาย",
	}
	db.Model(&Gender{}).Create(&Male)

	Female := Gender{
		Name: "หญิง",
	}
	db.Model(&Gender{}).Create(&Female)

	//---Status Data
	checkPayment := Status{
		Name: "ตรวจสอบการชำระเงิน",
	}
	db.Model(&Status{}).Create(&checkPayment)

	prepareForDelivery := Status{
		Name: "เตรียมการจัดส่ง ",
	}
	db.Model(&Status{}).Create(&prepareForDelivery)
	
	shipped := Status{
		Name: "จัดส่งแล้ว ",
	}
	db.Model(&Status{}).Create(&shipped)
}
