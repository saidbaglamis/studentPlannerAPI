package models

type Plan struct {
	ID         uint   `gorm:"primaryKey"`
	StudentID  uint   `gorm:"not null"`
	PlanInfo   string `gorm:"type:text;not null"`
	PlanStatus string `gorm:"type:text;not null"`
	StartDate  uint   `gorm:"check:number >= 0 AND number <= 2400"`
	EndDate    uint   `gorm:"check:number >= 0 AND number <= 2400"`
}
