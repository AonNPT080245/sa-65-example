package entity

import (
	"time"

	"gorm.io/gorm"
)

// type User struct {
// 	gorm.Model

// 	FirstName string

// 	LastName string

// 	Email string

// 	Age uint8

// 	BirthDay time.Time
// }

type Admin struct {
	gorm.Model
	User_name string `gorm:"uniqueIndex"`
	Password  string
	Name      string
	Employee  []Employee `gorm:"foreignKey:AdminID"`
}

type Department struct {
	gorm.Model
	DeptName    string
	Description string
	Get_staff   bool
	Employee    []Employee `gorm:"foreignKey:DepartmentID"`
	Salary      []Salary   `gorm:"foreignKey:DepartmentID"`
}

type Position struct {
	gorm.Model
	Name        string
	Description string
	Employee    []Employee `gorm:"foreignKey:PositionID"`
	Salary      []Salary   `gorm:"foreignKey:PositionID"`
}

type Salary struct {
	gorm.Model
	DepartmentID *uint
	Department   Department
	PositionID   *uint
	Position     Position
	Amount       int
	Employee     []Employee `gorm:"foreignKey:SalaryID"`
}

type Employee struct {
	gorm.Model
	Name    string
	Gender  string
	Age     uint
	Contact string
	Address string
	Date    time.Time

	AdminID *uint
	Admin   Admin

	DepartmentID *uint
	Department   Department

	PositionID *uint
	Position   Position

	SalaryID *uint
	Salary   Salary
}
