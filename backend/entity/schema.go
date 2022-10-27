package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"uniqueIndex"`
	Password string `json:"-"`
	Role     string
	Symptoms []Symptom `gorm:"foreignKey:DoctorID"`
}

type Doctor struct {
	gorm.Model
	Name  string
	Email string `gorm:"unigueIndex"`
	//  Password    string

	Symptoms []Symptom `gorm:"foreignKey:DoctorID"`
}

type Patient struct {
	gorm.Model
	Name    string
	Email   string `gorm:"unigueIndex"`
	Blood   string
	Gender  string
	Disease string

	Symptoms []Symptom `gorm:"foreignKey:PatientID"`
}

type Department struct {
	gorm.Model
	Name string

	Symptoms []Symptom `gorm:"foreignKey:DepartmentID"`
}

type Tenderness struct {
	gorm.Model
	Name string

	Symptoms []Symptom `gorm:"foreignKey:TendernessID"`
}

type Symptom struct {
	gorm.Model
	Explain     string
	SymptomTime time.Time

	// DoctorID เป็น FK
	UserID *uint
	// ข้อมูลของ Doctor เมื่อ join ตาราง
	User User

	// DoctorID เป็น FK
	DoctorID *uint
	// ข้อมูลของ Doctor เมื่อ join ตาราง
	Doctor Doctor

	// PatientID เป็น FK
	PatientID *uint
	// ข้อมูลของ Patient เมื่อ join ตาราง
	Patient Patient

	// TendernessID เป็น FK
	TendernessID *uint
	// ข้อมูลของ Tenderness เมื่อ join ตาราง
	Tenderness Tenderness

	// DepartmentID เป็น FK
	DepartmentID *uint
	// ข้อมูลของ Department เมื่อ join ตาราง
	Department Department
}
