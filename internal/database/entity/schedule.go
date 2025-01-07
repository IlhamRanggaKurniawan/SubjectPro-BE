package entity

import "time"

type Schedule struct {
	Id        uint64    `gorm:"primaryKey;autoIncrement"`
	Day       string    `gorm:"not null"`
	StartTime string    `gorm:"not null"`
	EndTime   string    `gorm:"not null"`
	SubjectId uint64    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
