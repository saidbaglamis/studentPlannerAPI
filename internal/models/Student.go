package models

type Student struct {
	ID           uint   `gorm:"primaryKey"`
	Name         string `gorm:"size:255;not null"`
	Email        string `gorm:"size:255;not null"`
	SchoolNumber string `gorm:"size:50;not null"`
	Class        string `gorm:"size:50;not null"`
	Plans        []Plan `gorm:"foreignKey:StudentID"`
}
