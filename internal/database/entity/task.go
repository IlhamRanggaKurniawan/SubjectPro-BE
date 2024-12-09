package entity

import "time"

type Task struct {
	Id        uint64    `gorm:"primaryKey;autoIncrement"`
	Type      string    `gorm:"not null"`
	Note      string    `gorm:"type:text"`
	Deadline  time.Time `gorm:"not null"`
	SubjectId uint64    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
