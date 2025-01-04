package entity

import "time"

type Schedule struct {
	Id        uint64    `gorm:"primaryKey;autoIncrement"`
	Day       string    `gorm:"not null"`
	StartTime time.Time `gorm:"not null"`
	EndTime   time.Time `gorm:"not null"`
	SubjectId uint64    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
