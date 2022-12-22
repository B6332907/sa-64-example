package entity

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("sa-64.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	database.AutoMigrate(
		&Role{},
		&Prefix{},
		&Policing{},
		&Gender{},
		&Patiend{},
		&Officer{},
	)

	db = database

	//Prefix data

	prefix_one := Prefix{
		Description: "นาย",
	}
	db.Model(&Prefix{}).Create(&prefix_one)

	prefix_two := Prefix{
		Description: "นาง",
	}
	db.Model(&Prefix{}).Create(&prefix_two)

	prefix_three := Prefix{
		Description: "นางสาว",
	}
	db.Model(&Prefix{}).Create(&prefix_three)

	// gender data
	gender_one := Gender{
		Description: "ชาย",
	}
	db.Model(&Gender{}).Create(&gender_one)

	gender_two := Prefix{
		Description: "หญิง",
	}
	db.Model(&Gender{}).Create(&gender_two)
	//role data
	role_one := Role{
		Description: "เจ้าหน้าที่คัดกรองคนไข้",
	}
	db.Model(&Role{}).Create(&role_one)

	role_two := Role{
		Description: "เจ้าหน้าที่ฝ่ายจัดการคนไข้",
	}
	db.Model(&Role{}).Create(&role_two)

	role_three := Role{
		Description: "เจ้าหน้าที่ฝ่ายรถฉุกเฉิน",
	}
	db.Model(&Role{}).Create(&role_three)

	//Policing data
	policing_one := Policing{
		Description: "สิทธิ์ข้าราชการ",
	}
	db.Model(&Policing{}).Create(&policing_one)

	policing_two := Policing{
		Description: "สิทธิ์บัตรทอง",
	}
	db.Model(&Policing{}).Create(&policing_two)

	policing_three := Policing{
		Description: "สิทธิ์ประกันสังคม",
	}
	db.Model(&Policing{}).Create(&policing_three)

	policing_four := Policing{
		Description: "สิทธิ์องกรคู่สัญญา",
	}
	db.Model(&Policing{}).Create(&policing_four)

	policing_five := Policing{
		Description: "ผู้ป่วยถือบัตรประกันสุขภาพ",
	}
	db.Model(&Policing{}).Create(&policing_five)

	policing_six := Policing{
		Description: "ไม่มี",
	}
	db.Model(&Policing{}).Create(&policing_six)

}
