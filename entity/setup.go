package entity

import (
	"fmt"
	"time"

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

	//database.AutoMigrate(&User{})
	database.AutoMigrate(
		&Admin{},
		&Department{},
		&Position{},
		&Salary{},
		&Employee{},
	)

	db = database

	// -------------------(Create value User_Admin)-----------------------------------
	db.Model(&Admin{}).Create(&Admin{
		User_name: "admin01",
		Password:  "Admin154",
		Name:      "Jardang guitar",
	})
	db.Model(&Admin{}).Create(&Admin{
		User_name: "admin02",
		Name:      "Bigtu",
		Password:  "Admin124",
	})

	// -------------------(ค้นหา AdminID ด้วย User_name ที่เข้าระบบมาใส่ใน  Entity Employee)-------------------------------
	var Jardang Admin
	var Bigtu Admin
	db.Raw("SELECT * FROM admins WHERE user_name = ?", "admin01").Scan(&Jardang)
	db.Raw("SELECT * FROM admins WHERE user_name = ?", "admin02").Scan(&Bigtu)

	// -------------------(Create value Position)-----------------------------------

	post1 := Position{
		Name:        "Manager",
		Description: "ผู้จัดการ...",
	}
	db.Model(&Position{}).Create(&post1)

	post2 := Position{
		Name:        "Head",
		Description: "หัวหน้าแผนก...",
	}
	db.Model(&Position{}).Create(&post2)

	post3 := Position{
		Name:        "Officer",
		Description: "เจ้าหน้าที่...",
	}
	db.Model(&Position{}).Create(&post3)

	post4 := Position{
		Name:        "general staff",
		Description: "พนักงานทั่วไป",
	}
	db.Model(&Position{}).Create(&post4)

	// -------------------(Create value Department)-----------------------------------

	dept1 := Department{
		DeptName:    "Financial Department",
		Description: "แผนกการเงิน",
		Get_staff:   true,
	}
	db.Model(&Department{}).Create(&dept1)

	dept2 := Department{
		DeptName:    "Security Department",
		Description: "แผนกฝ่ายรักษาความปลอดภัย",
		Get_staff:   true,
	}
	db.Model(&Department{}).Create(&dept2)

	dept3 := Department{
		DeptName:    "Repair Department",
		Description: "แผนกซ่อมแวมสำหรับติดต่อแจ้งซ่อม",
		Get_staff:   true,
	}
	db.Model(&Department{}).Create(&dept3)

	// -------------------(Create value Salary)-----------------------------------

	sal1 := Salary{
		Department: dept1,
		Position:   post2,
		Amount:     52000,
	}
	db.Model(&Salary{}).Create(&sal1)

	sal2 := Salary{
		Department: dept2,
		Position:   post3,
		Amount:     2800,
	}
	db.Model(&Salary{}).Create(&sal2)

	sal3 := Salary{
		Department: dept3,
		Position:   post2,
		Amount:     2100,
	}
	db.Model(&Salary{}).Create(&sal3)

	// -------------------(Create value Employee)-----------------------------------

	db.Model(&Employee{}).Create(&Employee{
		Department: dept1,
		Position:   post2,
		Salary:     sal1,
		Admin:      Bigtu,
		Name:       "พระบิดา",
		Gender:     "Male",
		Age:        62,
		Contact:    "084211xxx",
		Address:    "90/3 บ.หนองน้ำ",
		Date:       time.Now(),
	})

	db.Model(&Employee{}).Create(&Employee{
		Department: dept2,
		Position:   post3,
		Salary:     sal2,
		Admin:      Jardang,
		Name:       "จารย์แดง กีต้าไฟ",
		Gender:     "Male",
		Age:        64,
		Contact:    "089233xxx",
		Address:    "111/7 บ.หนองฮี",
		Date:       time.Now(),
	})

	db.Model(&Employee{}).Create(&Employee{
		Department: dept3,
		Position:   post2,
		Salary:     sal3,
		Admin:      Bigtu,
		Name:       "ลุงโทนี่",
		Gender:     "Male",
		Age:        52,
		Contact:    "089111xxx",
		Address:    "12/4 บ.บึง",
		Date:       time.Now(),
	})

	// -------------------(Querystep)-----------------------------------

	var target Admin
	db.Model(&Admin{}).Find(&target, db.Where("user_name = ?", "admin02"))

	var employeeList []*Employee
	db.Model(&Employee{}).
		Joins("Admin").
		Joins("Position").
		Joins("Department").
		Joins("Salary").
		Find(&employeeList, db.Where("admin_id = ?", target.ID))

	for _, wl := range employeeList {
		fmt.Printf("Employee: %v \n", wl.ID)
		fmt.Printf("Employee name: %v \n", wl.Name)
		fmt.Printf("gender: %v \n", wl.Gender)
		fmt.Printf("Age: %v \n", wl.Age)
		fmt.Printf("Contact: %v \n", wl.Contact)
		fmt.Printf("Address: %v \n", wl.Address)
		fmt.Printf("Date in: %v \n", wl.Date)
		fmt.Printf("Admin name: %v \n", wl.Admin.Name)
		fmt.Printf("staff position: %v \n", wl.Position.Name)
		fmt.Printf("Department: %v \n", wl.Department.DeptName)
		fmt.Printf("Salary amount: %v \n", wl.Salary.Amount)
		fmt.Printf("==========================\n")
	}

}
