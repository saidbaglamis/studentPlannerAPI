package models

type Student struct {
	ID           uint   `gorm:"primaryKey"`
	Email        string `gorm:"size:255;not null"`
	SchoolNumber string `gorm:"size:50;not null"`
	Class        string `gorm:"size:50;not null"`
	Plans        []Plan `gorm:"foreignKey:StudentID"`
}
