package models

import "time"

type Plan struct {
	ID         uint      `gorm:"primaryKey"`
	StudentID  uint      `gorm:"not null"`
	PlanInfo   string    `gorm:"type:text;not null"`
	PlanStatus string    `gorm:"type:text;not null"`
	StartDate  time.Time `gorm:"column:start_date;not null"`
	EndDate    time.Time `gorm:"coloumn:end_date;not null"`
}
